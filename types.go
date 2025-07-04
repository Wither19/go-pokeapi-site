package main

import "github.com/mtslzr/pokeapi-go/structs"


type Genus struct{
	Genus string "json:\"genus\""; 
	Language struct{
		Name string "json:\"name\"";
		URL string "json:\"url\""
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
	Pokemon structs.Pokemon
	PaddedID string
	EnglishGenus string
	FlavorTexts []FlavorText
	DisplayVersions []string
}