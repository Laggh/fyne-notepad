package main

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func printLn(a ...any) {
	fmt.Println(a...)
}

func getEmptyIcon() fyne.Resource {
	sgvStr := `<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"></svg>`
	svgArr := []byte(sgvStr)
	emptyIcon := fyne.NewStaticResource("empty.svg", svgArr)
	return emptyIcon
}
