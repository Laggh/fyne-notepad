package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("AppLegal")
	w := a.NewWindow("Exemplo Submenu")

	state := NewAppState(a, w)

	w.SetContent(state.Editor)
	w.SetMainMenu(getMenuBar(state))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
