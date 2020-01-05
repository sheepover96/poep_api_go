package models

import (
	"time"
)

type PoemTheme struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Detail          string    `json:"detail"`
	AnswerLengthMin int       `json:"answer_length_min"`
	AnswerLengthMax int       `json:"answer_length_max"`
	ThemeSetterName string    `json:"theme_setter_name"`
	CreatedAt       time.Time `json:"created_at"`
}
