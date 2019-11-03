package models

type PoemTag struct {
	ID          int    `json:"id"`
	Tag         string `json:"tag"`
	PoemThemeId int    `json:"poem_theme_id"`
}
