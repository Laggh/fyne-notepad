package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type AppState struct {
	App    fyne.App
	Window fyne.Window
	Editor *widget.Entry

	FilePath    string
	madeChanges bool
}

func handleEditorChange(state *AppState) {
	printLn("hEditorChange")
	state.madeChanges = true
	updateWindowName(state)
}

func updateWindowName(state *AppState) error {
	printLn("updateWindowName", state.FilePath, state.madeChanges)
	if state.FilePath == "" {
		if state.madeChanges {
			state.Window.SetTitle("*Sem Titulo")
		} else {
			state.Window.SetTitle("Sem Titulo")
		}
		return nil
	}

	split := strings.Split(state.FilePath, "/")
	fileName := split[len(split)-1]
	if state.madeChanges {
		state.Window.SetTitle("*" + fileName)
	} else {
		state.Window.SetTitle(fileName)
	}

	return nil
}

func handleOpenFile(state *AppState, path string) error {
	printLn("hOpenFile", path)
	data, err := loadFileRaw(path)
	if err != nil {
		dialog.ShowError(err, state.Window)
		return err
	}

	state.Editor.SetText(data)
	state.FilePath = path
	state.madeChanges = false
	updateWindowName(state)

	return nil
}

func handleSaveFile(state *AppState, path string) error {
	printLn("hSaveFile")

	err := saveFileRaw(path, state.Editor.Text)
	if err != nil {
		dialog.ShowError(err, state.Window)
		return err
	}

	state.madeChanges = false
	dialog.ShowInformation("", "Salvo com sucesso", state.Window)
	updateWindowName(state)
	return nil
}

func NewAppState(a fyne.App, w fyne.Window) *AppState {
	//TODO(): Fazer tudo né, tipo durr
	state := &AppState{
		App:    a,
		Window: w,
		Editor: widget.NewMultiLineEntry(),

		FilePath:    "",
		madeChanges: false,
	}

	state.Editor.OnChanged = func(s string) {
		handleEditorChange(state)
	}
	updateWindowName(state)
	return state

}
