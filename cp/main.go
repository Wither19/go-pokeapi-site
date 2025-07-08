package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/samber/lo"
)

// Since the API wants to be difficult on my home PC, this function gets the directory of the desired resource in the api-data folder.
func getAPILink(cat string, id string) string {
	return fmt.Sprintf("../api-data/%v/%v/index.json", cat, id)
}

// Reads and unmarshals the json for a Pokemon's general API entry, as specified by [id]
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

func main() {
	var id string

	var lvl string

	var atk string
	var def string
	var sta string

	fmt.Print("Provide Pokedex #: ")
	fmt.Scanln(&id)

	fmt.Print("Provide level: ")
	fmt.Scanln(&lvl)

	intLvl, err := strconv.ParseInt(lvl, 0, 0)
	if (err != nil) {
		log.Fatalln("Number conversion error:", err)
	}

	fmt.Print("Provide attack IV: ")
	fmt.Scanln(&atk)
	
	fmt.Print("Provide defense IV: ")
	fmt.Scanln(&def)
	
	fmt.Print("Provide Stamina IV: ")
	fmt.Scanln(&sta)

	c := getPkmn(id)

	var response string
	fmt.Printf("Calculating for #%d %v at Level %v with IVs of %v, %v and %v. is this correct? (y / n) ", c.ID, c.Name, lvl, atk, def, sta)
	fmt.Scanln(&response)

	if (strings.ToLower(response) != "y") {
		os.Exit(1)
	}

	intAtk, err := strconv.ParseInt(atk, 0, 0)
	if (err != nil) {
		log.Fatalln("Number conversion error:", err)
	}

	intDef, err := strconv.ParseInt(atk, 0, 0)
	if (err != nil) {
		log.Fatalln("Number conversion error:", err)
	}

	intSta, err := strconv.ParseInt(atk, 0, 0)
	if (err != nil) {
		log.Fatalln("Number conversion error:", err)
	}

	statsObj := getPkmn(id).Stats

	stats := lo.Map(statsObj, func(t struct{BaseStat int "json:\"base_stat\""; Effort int "json:\"effort\""; Stat struct{Name string "json:\"name\""; URL string "json:\"url\""} "json:\"stat\""}, i int) int {
		return t.BaseStat
	})

	floatStats := lo.Map(stats, func(t int, i int) float64 {
		return float64(t)
	})	

	fmt.Println(CPCalc(floatStats, int(intLvl), int(intAtk), int(intDef), int(intSta)))

}