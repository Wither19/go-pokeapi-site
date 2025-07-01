package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"slices"

	"github.com/mtslzr/pokeapi-go/structs"
)

func parseTemp(f string) *template.Template {
	return template.Must(template.ParseFiles(f))
}

func getAPILink(cat string, id string) string {
	return fmt.Sprintf("api-data/%v/%v/index.json", cat, id)
}

func leadingZeroes(num int, length int) string {
	return fmt.Sprintf("%0*d", length, num)
}

func serverSassComp(verbose bool) {
	sassSource := "./static/scss/App.scss"
	newCss := "./static/css/style.css"

	sassBuild := exec.Command("sass", sassSource, newCss, "--no-source-map")
	if err := sassBuild.Run(); err != nil {
		log.Fatalln("Sass build error:", err)

	} else if (verbose) {
		fmt.Printf("Sass successfully transpiled to %v\n", newCss)

	}
}

func getNatlDex() structs.Pokedex {
	dexURL := getAPILink("pokedex", "1")

	dex, err := os.ReadFile(dexURL)
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

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	d := getNatlDex().PokemonEntries

	serverSassComp(false)

	parseTemp("main.html").Execute(w, d)
}

func pkmnLoadfunc(w http.ResponseWriter, r *http.Request) {
	pkmnID := r.PathValue("id")

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
	}

	pkmn := getPkmn(pkmnID)
	species := getPkmnSpecies(pkmnID)

	paddedID := leadingZeroes(pkmn.ID, 4)

	var engGenus Genus

	for _, genus := range species.Genera {
		if genus.Language.Name == "en" {
			engGenus = genus
			break
		}
	}

	flavorTexts := []FlavorText{}
	engFlavors := []FlavorText{}

	omissions := []string{
		"red", "blue", "yellow", "gold", "silver", "crystal",
		"ruby", "sapphire", "emerald", "firered", "leafgreen",
		"diamond", "pearl", "platinum", "heartgold", "soulsilver", "black", "white", "black-2", "white-2",
	}
	
	for _, flavorText := range species.FlavorTextEntries {
		if (flavorText.Language.Name == "en") {
			engFlavors = append(engFlavors, flavorText)
		}
	}

	for _, flavorText := range engFlavors {
		if (!slices.Contains(omissions, flavorText.Version.Name)) {
			flavorTexts = append(flavorTexts, flavorText)
		}
	}

	data := PkmnData{Pokemon: pkmn, PaddedID: paddedID, EnglishGenus: engGenus.Genus, FlavorTexts: flavorTexts}

	serverSassComp(false)

	parseTemp("pkmn.html").Execute(w, data)

}

func main() {
	serverSassComp(true)
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", mainPageHandle)
	http.HandleFunc("/pkmn/{id}", pkmnLoadfunc)

	serverPort := ":8080"
	fmt.Printf("Server active at %v\n", serverPort)

	http.ListenAndServe(serverPort, nil)
}