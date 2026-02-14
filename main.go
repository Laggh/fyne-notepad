package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.NewWithID("AppLegal")
	w := a.NewWindow("Exemplo Submenu")

	state := NewAppState(a, w)

	editorTheme := &EditorTheme{Theme: theme.DefaultTheme(), State: state}

	w.SetContent(container.NewThemeOverride(state.Editor, editorTheme))
	w.SetMainMenu(state.MenuBar)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
