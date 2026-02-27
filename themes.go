package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type EditorTheme struct {
	fyne.Theme //Tema base para herdar as outras propriedades
	State      *AppState
}

func (t *EditorTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return t.State.fontSize * t.State.zoom
	}

	return t.Theme.Size(name)
}

func (t *EditorTheme) Font(style fyne.TextStyle) fyne.Resource {
	if t.State.fontResource != nil {
		return t.State.fontResource
	}
	return t.Theme.Font(style)
}
