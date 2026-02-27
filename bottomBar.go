package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Preguiça de fazer funcionar
func getBottomBar(state *AppState) *fyne.Container {
	positionLabel := widget.NewLabel("Ln 1 Col 1")
	zoomLabel := widget.NewLabel("100%")

	cont := container.NewHBox(
		layout.NewSpacer(),
		widget.NewSeparator(),
		positionLabel,
		widget.NewSeparator(),
		zoomLabel,
		widget.NewSeparator(),
		widget.NewLabel("Windows CRLF"),
		widget.NewSeparator(),
		widget.NewLabel("UTF-8"),
	)

	return cont
}
