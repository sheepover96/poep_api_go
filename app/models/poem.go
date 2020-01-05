package models

import (
	"time"
)

type Poem struct {
	ID           int       `json:"id"`
	PoemThemeID  int       `json:"poemThemeId"`
	Nfav         int       `json:"nfav"`
	AnswererName string    `json:"answererName"`
	AnswerText   string    `json:"answerText"`
	CreatedAt    time.Time `json:"createdAt"`
}
