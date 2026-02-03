package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func getOpenItemFunc(state *AppState) func() {
	return func() {
		fileOpenDiag := dialog.NewFileOpen(
			//callback
			func(r fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowError(err, state.Window)
					return
				}
				if r == nil {
					return
				}

				uriPath := r.URI().Path()
				handleOpenFile(state, uriPath)
			},
			//window
			state.Window,
		)
		fileOpenDiag.Show()
	}
}

func getSaveItemFunc(state *AppState) func() {
	return func() {
		if state.FilePath == "" {
			saveAsFunc := getSaveAsItemFunc(state)
			saveAsFunc()
			return
		}

		handleSaveFile(state, state.FilePath)
	}
}

func getSaveAsItemFunc(state *AppState) func() {
	return func() {
		fileSaveDiag := dialog.NewFileSave(
			//callback
			func(w fyne.URIWriteCloser, err error) {
				if err != nil {
					dialog.ShowError(err, state.Window)
					return
				}
				if w == nil {
					return
				}

				uriPath := w.URI().Path()
				state.FilePath = uriPath
				handleSaveFile(state, uriPath)

			},
			//window
			state.Window,
		)
		fileSaveDiag.Show()
	}
}

func getMenuBar(state *AppState) *fyne.MainMenu {

	openItem := fyne.NewMenuItem("Abrir...", getOpenItemFunc(state))
	saveItem := fyne.NewMenuItem("Salvar", getSaveItemFunc(state))
	saveAsItem := fyne.NewMenuItem("Salvar Como", getSaveAsItemFunc(state))

	archiveMenu := fyne.NewMenu("Arquivo", openItem, saveItem, saveAsItem)

	return fyne.NewMainMenu(archiveMenu)
}
