package main

import (
	"log"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/samber/lo"
)

// Since the API wants to be difficult on my home PC, this function gets the directory of the desired resource in the api-data folder.
// func getAPILink(cat string, id string) string {
// 	return fmt.Sprintf("api-data/%v/%v/index.json", cat, id)
// }

// Reads and unmarshals the json for the National Pokedex.
func getNatlDex() []NationalDexEntry {
	// dex, err := os.ReadFile(getAPILink("pokedex", "1"))
	// if err != nil {
	// 	log.Fatalln("Dex fetch error:", err)
	// }

	// var pokedex Pokedex

	// dexUnpackErr := json.Unmarshal(dex, &pokedex)
	// if dexUnpackErr != nil {
	// 	log.Fatalln("Dex unpack error:", err)
	// }

	d, err := pokeapi.Pokedex("1")
	if err != nil {
		log.Fatalln("Dex fetch error:", err)
	}

	e := lo.Map(d.PokemonEntries, func(item struct {
		EntryNumber    int `json:"entry_number"`
		PokemonSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
	}, index int) NationalDexEntry {
		var new NationalDexEntry

		new.EntryNumber = item.EntryNumber
		new.PokemonSpecies = item.PokemonSpecies

		return new
	})

	return e
}

// Fetches a Pokemon's general API entry, as specified by [id]
func getPkmn(id string) structs.Pokemon {
	p, err := pokeapi.Pokemon(id)
	if err != nil {
		log.Fatalln("Pokemon fetch error:", err)
	}

	return p
}

// Reads and unmarshals the json for a Pokemon's species API entry, as specified by [id]
func getPkmnSpecies(id string) structs.PokemonSpecies {
	s, sErr := pokeapi.PokemonSpecies(id)
	if sErr != nil {
		log.Fatalln("Species fetch error:", sErr)
	}

	return s
}
