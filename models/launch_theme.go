package models

type LaunchTheme struct {
	PoemThemeLaunch *PoemTheme `json:"poem_theme"`
	InitialPoem     *Poem      `json:"poem"`
	Tags            []PoemTag  `json:"poem_tags"`
}
