package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func getMenuBar(state *AppState) *fyne.MainMenu {

	openItem := fyne.NewMenuItem("Abrir...", func() {
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
	})
	saveItem := fyne.NewMenuItem("Salvar", func() {
		if state.FilePath == "" {
			printLn("tentando salvar sem filepath")
			return
		}

		handleSaveFile(state, state.FilePath)
	})

	saveAsItem := fyne.NewMenuItem("Salvar Como", func() {
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
				handleSaveFile(state, uriPath)

			},
			//window
			state.Window,
		)
		fileSaveDiag.Show()
	})

	archiveMenu := fyne.NewMenu("Arquivo", openItem, saveItem, saveAsItem)

	return fyne.NewMainMenu(archiveMenu)
}
