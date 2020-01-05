package api

import (
	"net/http"
	// "fmt"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	db "github.com/sheepover96/poep_api/app/db"
	models "github.com/sheepover96/poep_api/app/models"
)

func PoemRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/poem/:id", GetPoemByID)
		v1.GET("/poem_theme/:id/poems", GetPoemsForPoemTheme)
		v1.GET("/poem_theme/:id/poems/:order", GetPoemsForPoemTheme)
	}

}

func GetPoemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	}
	rows, err := db.DBcon.Query(`select id, poem_theme_id, nfav, answerer_name, answer_text,
															date_created from poem where id = ?;`, id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	}

	poem := new(models.Poem)
	for rows.Next() {
		if err := rows.Scan(&poem.ID, &poem.PoemThemeID, &poem.Nfav, &poem.CreatedAt,
			&poem.AnswererName, &poem.AnswerText); err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	jsonBytes, err := json.Marshal(poem)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	c.JSON(http.StatusOK, string(jsonBytes))
}

func GetPoemsForPoemTheme(c *gin.Context) {
	order := c.Param("order")

	rows, err := db.DBcon.Query(`select id, poem_theme_id, nfav, answerer_name, answer_text,
															date_created from poem order by ?;`, order)
	if err != nil {
		log.Fatal(err)
	}

	var poemList []models.Poem
	for rows.Next() {
		var poem models.Poem
		if err := rows.Scan(&poem.ID, &poem.PoemThemeID, &poem.Nfav,
			&poem.AnswererName, &poem.AnswerText, &poem.CreatedAt); err != nil {
			log.Fatal(err)
		}
		poemList = append(poemList, poem)
	}

	jsonBytes, err := json.Marshal(poemList)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	c.JSON(http.StatusOK, string(jsonBytes))
}
