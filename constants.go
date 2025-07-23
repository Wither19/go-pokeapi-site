package main

var natlDexEntries []struct {
	EntryNumber    int "json:\"entry_number\""
	PokemonSpecies struct {
		Name string "json:\"name\""
		URL  string "json:\"url\""
	} "json:\"pokemon_species\""
} = getNatlDex().PokemonEntries

var omissions []string = []string{
	"red", "blue", "yellow", "gold", "silver", "crystal",
	"ruby", "sapphire", "emerald", "firered", "leafgreen",
	"diamond", "pearl", "platinum", "heartgold", "soulsilver", "black", "white", "black-2", "white-2",
}

var regionKeywords []string = []string{"kanto", "johto", "hoenn", "sinnoh", "unova", "kalos", "alola", "unknown", "galar", "hisui", "paldea"}
var regionStarts []int = []int{1, 152, 252, 387, 494, 650, 722, 808, 810, 899, 906}
