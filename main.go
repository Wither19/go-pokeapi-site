package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/*
var staticWebFiles embed.FS

func main() {
	staticContent, err := fs.Sub(staticWebFiles, "static")
	if err != nil {
		log.Fatalln("Could not embed static files:", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))

	http.HandleFunc("/", mainPageHandle)
	http.HandleFunc("/search/{search}", mainPagePkmnSearch)
	http.HandleFunc("/search/", pkmnSearchNotFound)

	http.HandleFunc("/pkmn/{id}", pkmnLoad)

	http.HandleFunc("/pkmn/{id}/prev", prevPkmnLoad)
	http.HandleFunc("/pkmn/{id}/next", nextPkmnLoad)

	http.HandleFunc("/random", pkmnRandomize)

	serverPort := ":8080"

	http.ListenAndServe(serverPort, nil)
}
