package main

import "github.com/samber/lo"

var natlDexEntries []NatlDexEntry = getNatlDex()

var omissions []string = []string{
	"red", "blue", "yellow", "gold", "silver", "crystal",
	"ruby", "sapphire", "emerald", "firered", "leafgreen",
	"diamond", "pearl", "platinum", "heartgold", "soulsilver", "black", "white", "black-2", "white-2",
}

var regionKeywords []string = []string{"kanto", "johto", "hoenn", "sinnoh", "unova", "kalos", "alola", "unknown", "galar", "hisui", "paldea"}
var regionStarts []int = []int{1, 152, 252, 387, 494, 650, 722, 808, 810, 899, 906}

var pkmnNames []string = lo.Map(natlDexEntries, func(item NatlDexEntry, _ int) string {
	return item.PokemonSpecies.Name
})
