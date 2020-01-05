package models

type PoemTag struct {
	ID          int    `db:id`
	Tag         string `db:tag`
	PoemThemeId int    `db:poem_theme_id`
}
