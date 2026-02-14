package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func baseFunc(state *AppState) func() {
	return func() {
		dialog.ShowInformation("Example", "Hello, World!", state.Window)
	}
}

// Arquivo
// TODO(): Melhorar o sistema de perguntar se pode salvar, e ter 3 botões ao invez de 2
func getNewItemFunc(state *AppState) func() {
	return func() {
		_f := func(b bool) {
			if b == false {
				return
			}
			state.FilePath = ""
			state.Editor.SetText("") //muda o madeChanges para true
			state.madeChanges = false
			updateWindowName(state)
		}

		if state.madeChanges {
			diag := dialog.NewConfirm("notepad",
				"Você tem mudanças não salvas, deseja salvar",
				_f, state.Window)
			diag.Show()
		} else {
			_f(true)
		}

	}
}

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

//Editar

// Formatar
func getLineBreakItemFunc(state *AppState) func() {
	return func() {
		lineBreakItem := state.MenuBar.Items[2].Items[0]
		if state.DoLineBreak {
			state.DoLineBreak = false
			lineBreakItem.Checked = false
			state.Editor.Wrapping = fyne.TextWrapOff
			state.Editor.Refresh() //atualiza para arrumar a quebra de linha
		} else {
			state.DoLineBreak = true
			lineBreakItem.Checked = true
			state.Editor.Wrapping = fyne.TextWrapBreak
			state.Editor.Refresh() //atualiza para arrumar a quebra de linha
		}

	}
}

func getFontItemFunc(state *AppState) func() {
	return func() {
		//fontCard
		fontCardContent := newScrolbarSelection()
		fontCard := widget.NewCard("Fonte", "", fontCardContent)

		//styleCard
		exampleEntry := widget.NewMultiLineEntry()
		exampleEntry.Text = "Exemplo \n\tTab"
		var boldCheck *widget.Check
		var italicCheck *widget.Check
		var monospaceCheck *widget.Check

		refreshStyleCard := func() {
			state.Editor.Refresh()
			exampleEntry.TextStyle = *state.TextStyle
			exampleEntry.Refresh()

			if state.TextStyle.Monospace {
				boldCheck.Disable()
				italicCheck.Disable()
			} else {
				boldCheck.Enable()
				italicCheck.Enable()
			}

		}
		boldCheck = widget.NewCheck("Negrito", func(b bool) {
			state.TextStyle.Bold = b
			refreshStyleCard()
		})
		italicCheck = widget.NewCheck("Italico", func(b bool) {
			state.TextStyle.Italic = b
			refreshStyleCard()
		})
		monospaceCheck = widget.NewCheck("Mono-espaçado", func(b bool) {
			state.TextStyle.Monospace = b
			refreshStyleCard()
		})

		boldCheck.Checked = state.TextStyle.Bold
		italicCheck.Checked = state.TextStyle.Italic
		monospaceCheck.Checked = state.TextStyle.Monospace
		refreshStyleCard()
		styleCardContent := container.NewVSplit(
			exampleEntry,
			container.NewVBox(boldCheck, italicCheck, monospaceCheck),
		)
		styleCardContent.Offset = 0.2

		styleCard := widget.NewCard("Estilo da fonte", "", styleCardContent)

		//sizeCard
		sizeCardContent := widget.NewMultiLineEntry()
		sizeCard := widget.NewCard("Tamanho: ", "", sizeCardContent)

		content := container.NewGridWithColumns(3, fontCard, styleCard, sizeCard)

		diag := dialog.NewCustom("Fonte", "Fechar", content, state.Window)
		diag.Resize(fyne.NewSize(800, 600))
		diag.Show()
	}
}

func getMenuBar(state *AppState) *fyne.MainMenu {

	//Arquivo [0]
	newItem := fyne.NewMenuItem("Novo", getNewItemFunc(state))
	//TODO(): fazer janela novo funfar
	newWindowItem := fyne.NewMenuItem("*Nova Janela", baseFunc(state))
	openItem := fyne.NewMenuItem("Abrir...", getOpenItemFunc(state))
	saveItem := fyne.NewMenuItem("Salvar", getSaveItemFunc(state))
	saveAsItem := fyne.NewMenuItem("Salvar Como", getSaveAsItemFunc(state))
	archiveMenu := fyne.NewMenu("Arquivo", newItem, newWindowItem, openItem, saveItem, saveAsItem)

	//Editar [1]
	blankItem := fyne.NewMenuItem("Nada", baseFunc(state))
	editMenu := fyne.NewMenu("Editar", blankItem)

	//Formatar [2]
	//lineBreakItem := fyne.NewMenuItemWithIcon("Quebra de Linha Automatica", theme.AccountIcon(), getLineBreakItemFunc(state))
	lineBreakItem := fyne.NewMenuItem("Quebra de Linha Automatica", getLineBreakItemFunc(state))
	lineBreakItem.Icon = getEmptyIcon()
	fontItem := fyne.NewMenuItem("Fonte", getFontItemFunc(state))
	formatMenu := fyne.NewMenu("Formatar", lineBreakItem, fontItem)

	//Exibir [3]
	zoomItem := fyne.NewMenuItem("Zoom", baseFunc(state))
	statusBarItem := fyne.NewMenuItem("Barra de Status", baseFunc(state))
	displayMenu := fyne.NewMenu("Exibir", zoomItem, statusBarItem)

	//Ajuda [4]
	showHelpItem := fyne.NewMenuItem("Exibir ajuda", baseFunc(state))
	sendCommentItem := fyne.NewMenuItem("Enviar Comentario", baseFunc(state))
	aboutItem := fyne.NewMenuItem("Sobre o Bloco de notas", baseFunc(state))
	helpMenu := fyne.NewMenu("Exibir", showHelpItem, sendCommentItem, fyne.NewMenuItemSeparator(), aboutItem)

	return fyne.NewMainMenu(archiveMenu, editMenu, formatMenu, displayMenu, helpMenu)
}
