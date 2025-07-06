package main

import (
	"fmt"
	"html/template"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	"github.com/savioxavier/termlink"
)

func parseTemp(n string, f template.FuncMap) *template.Template {
	if (f != nil) {
		return template.Must(template.New(n).Funcs(sprig.FuncMap()).Funcs(f).ParseFiles(n))	
	} else {
		return template.Must(template.New(n).Funcs(sprig.FuncMap()).ParseFiles(n))
	}
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

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}