package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func baseFunc(state *AppState) func() {
	return func() {
		dialog.ShowInformation("Example", "Hello, World!", state.Window)
	}
}

// Arquivo
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

	//Arquivo
	newItem := fyne.NewMenuItem("Novo", baseFunc(state))
	newWindowItem := fyne.NewMenuItem("Nova Janela", baseFunc(state))
	openItem := fyne.NewMenuItem("Abrir...", getOpenItemFunc(state))
	saveItem := fyne.NewMenuItem("Salvar", getSaveItemFunc(state))
	saveAsItem := fyne.NewMenuItem("Salvar Como", getSaveAsItemFunc(state))
	archiveMenu := fyne.NewMenu("Arquivo", newItem, newWindowItem, openItem, saveItem, saveAsItem)

	//Editar
	blankItem := fyne.NewMenuItem("Nada", baseFunc(state))
	editMenu := fyne.NewMenu("Editar", blankItem)

	//Formatar
	lineBreakItem := fyne.NewMenuItem("Quebra de Linha Automatica", baseFunc(state))
	fontItem := fyne.NewMenuItem("Fonte", baseFunc(state))
	formatMenu := fyne.NewMenu("Formatar", lineBreakItem, fontItem)

	//Exibir
	zoomItem := fyne.NewMenuItem("Zoom", baseFunc(state))
	statusBarItem := fyne.NewMenuItem("Barra de Status", baseFunc(state))
	displayMenu := fyne.NewMenu("Exibir", zoomItem, statusBarItem)

	//Ajuda
	showHelpItem := fyne.NewMenuItem("Exibir ajuda", baseFunc(state))
	sendCommentItem := fyne.NewMenuItem("Enviar Comentario", baseFunc(state))
	aboutItem := fyne.NewMenuItem("Sobre o Bloco de notas", baseFunc(state))
	helpMenu := fyne.NewMenu("Exibir", showHelpItem, sendCommentItem, fyne.NewMenuItemSeparator(), aboutItem)

	return fyne.NewMainMenu(archiveMenu, editMenu, formatMenu, displayMenu, helpMenu)
}
