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

	// container.New com sizeLimiterLayout "mente" o tamanho mínimo,
	// impedindo a janela de explodir, mas entrega todo o espaço pro editor.
	w.SetContent(container.New(&sizeLimiterLayout{},
		container.NewThemeOverride(state.Editor, editorTheme)))

	w.SetMainMenu(state.MenuBar)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

// Layout simples para limitar o tamanho minimo reportado
type sizeLimiterLayout struct{}

func (l *sizeLimiterLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, o := range objects {
		o.Resize(size)
		o.Move(fyne.NewPos(0, 0))
	}
}

func (l *sizeLimiterLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(100, 100)
}
