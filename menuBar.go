package main

import (
	"strconv"

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
		//Precreating Vars
		exampleLabel := widget.NewLabel("AaBbYyZz")
		exampleLabelContainer := container.NewThemeOverride(exampleLabel, state.Editor.Theme())

		zoomLabel := widget.NewLabel("Zoom: Indefinido")

		//fontCard
		fontCardContent := container.NewCenter(widget.NewLabel("Em breve..."))
		fontCard := widget.NewCard("Fonte", "", fontCardContent)

		//styleCard
		var boldCheck *widget.Check
		var italicCheck *widget.Check
		var monospaceCheck *widget.Check

		refresh := func() {
			state.Editor.Refresh()

			if state.TextStyle.Monospace {
				boldCheck.Disable()
				italicCheck.Disable()
			} else {
				boldCheck.Enable()
				italicCheck.Enable()
			}

			exampleLabel.TextStyle = *state.TextStyle
			exampleLabel.Refresh()

			zoomLabel.SetText("Zoom: " + strconv.Itoa(int(state.zoom*100)) + "%")
			zoomLabel.Refresh()
		}
		boldCheck = widget.NewCheck("Negrito", func(b bool) {
			state.TextStyle.Bold = b
			refresh()
		})
		italicCheck = widget.NewCheck("Italico", func(b bool) {
			state.TextStyle.Italic = b
			refresh()
		})
		monospaceCheck = widget.NewCheck("Mono-espaçado", func(b bool) {
			state.TextStyle.Monospace = b
			refresh()
		})

		boldCheck.Checked = state.TextStyle.Bold
		italicCheck.Checked = state.TextStyle.Italic
		monospaceCheck.Checked = state.TextStyle.Monospace
		refresh()
		styleCardContent := container.NewVBox(boldCheck, italicCheck, monospaceCheck)

		styleCard := widget.NewCard("Estilo da fonte", "", styleCardContent)

		//sizeCard
		sizesStr := []string{"8", "9", "10", "11", "12", "14", "16", "18", "20", "22", "24", "26", "28", "36", "48", "72"}
		sizeEntry := widget.NewSelectEntry(sizesStr)
		//zoomLabel = widget.NewLabel("Zoom: 100%")
		zoomSlider := widget.NewSlider(0.1, 5)
		zoomSlider.Step = 0.1

		sizeCardContent :=
			container.NewVBox(
				sizeEntry,
				zoomLabel,
				zoomSlider,
			)

		sizeEntry.OnChanged = func(s string) {
			size, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return
			}
			if size > 200 {
				size = 200
			}
			if size < 1 {
				size = 1
			}

			state.fontSize = float32(size)
			refresh()
		}
		sizeEntry.SetText(strconv.Itoa(int(state.fontSize)))

		zoomSlider.OnChangeEnded = func(f float64) {
			state.zoom = float32(f)
			refresh()
		}
		zoomSlider.Value = float64(state.zoom)

		sizeCard := widget.NewCard("Tamanho", "", container.NewVBox(sizeCardContent))

		topControls := container.NewVBox(
			container.NewGridWithColumns(3, fontCard, styleCard, sizeCard),
			widget.NewSeparator(),
		)
		exampleCard := widget.NewCard("Exemplo", "", container.NewCenter(exampleLabelContainer))

		content := container.NewBorder(topControls, nil, nil, nil, exampleCard)

		diag := dialog.NewCustom("Fonte", "Fechar", content, state.Window)
		diag.Resize(fyne.NewSize(800, 600))
		diag.Show()
	}
}

// Exibir
func getZoomPlusItemFunc(state *AppState) func() {
	return func() {
		state.zoom += 0.1
		state.Editor.Refresh()
	}
}

func getZoomMinusItemFunc(state *AppState) func() {
	return func() {
		state.zoom -= 0.1
		state.Editor.Refresh()
	}
}

func getZoomDefaultItemFunc(state *AppState) func() {
	return func() {
		state.zoom = 1
		state.Editor.Refresh()
	}
}

// getMenuBar
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
	zoomPlusItem := fyne.NewMenuItem("Ampliar", getZoomPlusItemFunc(state))
	zoomMinusItem := fyne.NewMenuItem("Reduzir", getZoomMinusItemFunc(state))
	zoomDefaultItem := fyne.NewMenuItem("Restaurar Zoom Padrão", getZoomDefaultItemFunc(state))
	zoomItem.ChildMenu = fyne.NewMenu("", zoomPlusItem, zoomMinusItem, zoomDefaultItem)

	statusBarItem := fyne.NewMenuItem("Barra de Status", baseFunc(state))
	displayMenu := fyne.NewMenu("Exibir", zoomItem, statusBarItem)

	//Ajuda [4]
	showHelpItem := fyne.NewMenuItem("Exibir ajuda", baseFunc(state))
	sendCommentItem := fyne.NewMenuItem("Enviar Comentario", baseFunc(state))
	aboutItem := fyne.NewMenuItem("Sobre o Bloco de notas", baseFunc(state))
	helpMenu := fyne.NewMenu("Exibir", showHelpItem, sendCommentItem, fyne.NewMenuItemSeparator(), aboutItem)

	return fyne.NewMainMenu(archiveMenu, editMenu, formatMenu, displayMenu, helpMenu)
}
