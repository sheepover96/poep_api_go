package models

import (
	"time"
)

type Poem struct {
	ID           int       `db:"id"`
	PoemThemeID  int       `db:"poem_theme_id"`
	Nfav         int       `db:"nfav"`
	AnswererName string    `db:"answerer_name"`
	AnswerText   string    `db:"answer_text"`
	CreatedAt    time.Time `db:"created_at"`
}
