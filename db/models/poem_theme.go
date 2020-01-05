package models

import (
	"time"
)

type PoemTheme struct {
	ID              int       `db:id`
	Title           string    `db:title`
	Detail          string    `db:detail`
	AnswerLengthMin int       `db:answer_length_min`
	AnswerLengthMax int       `db:answer_length_max`
	ThemeSetterName string    `db:theme_setter_name`
	CreatedAt       time.Time `db:created_at`
}
