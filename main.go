package main

import (
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

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
