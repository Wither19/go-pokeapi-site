package main

import "github.com/mtslzr/pokeapi-go/structs"

type Pokedex struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID           int    `json:"id"`
	IsMainSeries bool   `json:"is_main_series"`
	Name         string `json:"name"`
	Names        []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEntries []NatlDexEntry `json:"pokemon_entries"`
	Region         interface{}    `json:"region"`
	VersionGroups  []interface{}  `json:"version_groups"`
}

type NatlDexEntry struct {
	EntryNumber    int `json:"entry_number"`
	PokemonSpecies struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type Genus struct {
	Genus    string "json:\"genus\""
	Language struct {
		Name string "json:\"name\""
		URL  string "json:\"url\""
	} "json:\"language\""
}

type FlavorText struct {
	FlavorText string `json:"flavor_text"`
	Language   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"language"`
	Version struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version"`
}

type PkmnData struct {
	Pokemon      structs.Pokemon
	PaddedID     string
	EnglishGenus string
	FlavorTexts  []FlavorText
	Config       ServerConfig
}

type ServerConfig struct {
	PokemonArtwork string `yaml:"pokemon-artwork"`
}
