package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

// Since the API wants to be difficult on my home PC, this function gets the directory of the desired resource in the api-data folder.
func getAPILink(cat string, id string) string {
	return fmt.Sprintf("api-data/%v/%v/index.json", cat, id)
}

// Reads and unmarshals the json for the National Pokedex.
func getNatlDex() []NationalDexEntry {
	dex, err := os.ReadFile(getAPILink("pokedex", "1"))
	if err != nil {
		log.Fatalln("Dex fetch error:", err)
	}

	var pokedex Pokedex

	dexUnpackErr := json.Unmarshal(dex, &pokedex)
	if dexUnpackErr != nil {
		log.Fatalln("Dex unpack error:", err)
	}

	return pokedex.PokemonEntries
}

// Reads and unmarshals the json for a Pokemon's general API entry, as specified by [id]
func getPkmn(id string) structs.Pokemon {
	pURL := getAPILink("pokemon", id)

	p, pErr := os.ReadFile(pURL)
	if pErr != nil {
		log.Fatalln("Pokemon fetch error:", pErr)
	}

	var pkmn structs.Pokemon

	pUnpackErr := json.Unmarshal(p, &pkmn)
	if pUnpackErr != nil {
		log.Fatalln("Pokemon unpack error:", pUnpackErr)
	}
	return pkmn
}

// Reads and unmarshals the json for a Pokemon's species API entry, as specified by [id]
func getPkmnSpecies(id string) structs.PokemonSpecies {
	sURL := getAPILink("pokemon-species", id)

	s, sErr := os.ReadFile(sURL)
	if sErr != nil {
		log.Fatalln("Species fetch error:", sErr)
	}

	var species structs.PokemonSpecies

	sUnpackErr := json.Unmarshal(s, &species)
	if sUnpackErr != nil {
		log.Fatalln("Species unpack error:", sUnpackErr)
	}
	return species
}
