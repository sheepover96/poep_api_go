package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	db "github.com/sheepover96/poep_api/app/db"
	models "github.com/sheepover96/poep_api/app/models"
)

func PoemThemeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/poem_theme/:id", GetPoemThemeByID)
		v1.POST("/compe/launch", LaunchPoemCompe)
	}

}

func GetPoemThemeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	}
	rows, err := db.DBcon.Query(`select id, title, detail, answer_length_min, answer_length_max,
															theme_setter_name, created_at from poem_theme where id = ?;`, id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
	}

	poemTheme := new(models.PoemTheme)
	for rows.Next() {
		if err := rows.Scan(&poemTheme.ID, &poemTheme.Title, &poemTheme.Detail,
			&poemTheme.AnswerLengthMin, &poemTheme.AnswerLengthMax, &poemTheme.ThemeSetterName,
			&poemTheme.CreatedAt); err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	jsonBytes, err := json.Marshal(poemTheme)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	c.JSON(http.StatusOK, string(jsonBytes))
}

func LaunchPoemCompe(c *gin.Context) {
	var poemCompe models.LaunchTheme
	if err := c.ShouldBindJSON(&poemCompe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	poemTheme := poemCompe.PoemThemeLaunch
	initialPoem := poemCompe.InitialPoem
	Tags := poemCompe.Tags
	if poemTheme == nil || initialPoem == nil || len(Tags) == 0 {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	err := db.Transact(db.DBcon, func(tx *sql.Tx) error {
		poemThemeIns, err := tx.Prepare(`insert into poem_theme(title, ntag, detail, npoem,
			answer_length_min, answer_length_max, theme_setter_name) values(?, ?, ?, ?, ?, ?, ?)`)
		if err != nil {
			return err
		}

		res, err := poemThemeIns.Exec(poemTheme.Title, 0, poemTheme.Detail, 0, poemTheme.AnswerLengthMin,
			poemTheme.AnswerLengthMax, poemTheme.ThemeSetterName)
		if err != nil {
			return err
		}

		poemThemeID, err := res.LastInsertId()
		if err != nil {
			return err
		}

		initialPoemIns, err := tx.Prepare(`insert into poem(poem_theme_id, nfav,
			answerer_name, answer_text) values(?, ?, ?, ?)`)
		if err != nil {
			return err
		}

		_, err = initialPoemIns.Exec(poemThemeID, 0, initialPoem.AnswererName,
			initialPoem.AnswerText)
		if err != nil {
			return err
		}

		sqlStr := "INSERT INTO poem_tag (tag, poem_theme_id) VALUES "
		vals := []interface{}{}
		for _, row := range Tags {
			sqlStr += "(?, ?),"
			vals = append(vals, row.Tag, poemThemeID)
		}
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		tagsIns, err := tx.Prepare(sqlStr)
		if err != nil {
			return err
		}

		if _, err := tagsIns.Exec(vals...); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	} else {
		c.String(http.StatusOK, "OK")
	}

}
