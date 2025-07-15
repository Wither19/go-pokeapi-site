package main

import (
	"net/http"
)

func main() {
	serverSassComp(true)
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", mainPageHandle)
	http.HandleFunc("/pkmn/{id}", pkmnLoad)

	http.HandleFunc("/pkmn/{id}/prev", prevPkmnLoad)
	http.HandleFunc("/pkmn/{id}/next", nextPkmnLoad)

	http.HandleFunc("/random", pkmnRandomize)

	serverPort := ":8080"
	serverActiveMsg("Server active at %v\n", serverPort)

	http.ListenAndServe(serverPort, nil)
}