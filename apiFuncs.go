package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mtslzr/pokeapi-go/structs"
)

func getAPILink(cat string, id string) string {
	return fmt.Sprintf("api-data/%v/%v/index.json", cat, id)
}

func getNatlDex() structs.Pokedex {
	dex, err := os.ReadFile(getAPILink("pokedex", "1"))
	if (err != nil) {
		log.Fatalln("Dex fetch error:", err)
	}

	var pokedex structs.Pokedex

	dexUnpackErr := json.Unmarshal(dex, &pokedex)
	if (dexUnpackErr != nil) {
		log.Fatalln("Dex unpack error:", err)
	}

	return pokedex
}

func getPkmn(id string) structs.Pokemon {
	pURL := getAPILink("pokemon", id)

	p, pErr := os.ReadFile(pURL)
	if (pErr != nil) {
		log.Fatalln("Pokemon fetch error:", pErr)
	}

	var pkmn structs.Pokemon

	pUnpackErr := json.Unmarshal(p, &pkmn)
	if (pUnpackErr != nil) {
		log.Fatalln("Pokemon unpack error:", pUnpackErr)
	}
	return pkmn
}

func getPkmnSpecies(id string) structs.PokemonSpecies {
	sURL := getAPILink("pokemon-species", id)
	
	s, sErr := os.ReadFile(sURL)
	if (sErr != nil) {
		log.Fatalln("Species fetch error:", sErr)
	}

	var species structs.PokemonSpecies

	sUnpackErr := json.Unmarshal(s, &species)
	if (sUnpackErr != nil) {
		log.Fatalln("Species unpack error:", sUnpackErr)
	}	
	return species
}