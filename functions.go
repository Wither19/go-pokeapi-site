package main

import (
	"fmt"
	"html/template"

	"github.com/savioxavier/termlink"
)

func parseTemp(f string) *template.Template {
	return template.Must(template.ParseFiles(f))
}

func getAPILink(cat string, id string) string {
	return fmt.Sprintf("api-data/%v/%v/index.json", cat, id)
}

func leadingZeroes(num int, length int) string {
	return fmt.Sprintf("%0*d", length, num)
}

func serverActiveMsg(msg string, port string) {
	fmt.Printf(msg, termlink.ColorLink(fmt.Sprintf("port %v", port), fmt.Sprintf("http://localhost%v", port), "blue"))
}