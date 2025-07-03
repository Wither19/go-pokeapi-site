package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"slices"

	"github.com/samber/lo"
)

func serverSassComp(l bool) {
	sassSource := "./static/scss/App.scss"
	newCss := "./static/css/style.css"

	sassBuild := exec.Command("sass", sassSource, newCss, "--no-source-map")
	if err := sassBuild.Run(); err != nil {
		log.Fatalln("Sass build error:", err)

	} else if (l) {
		fmt.Printf("Sass successfully transpiled to %v\n", newCss)

	}
}

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	d := getNatlDex().PokemonEntries

	serverSassComp(false)

	parseTemp("main.tmpl").Execute(w, d)
}

func pkmnLoadfunc(w http.ResponseWriter, r *http.Request) {
	pkmnID := r.PathValue("id")

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

	omissions := []string{
		"red", "blue", "yellow", "gold", "silver", "crystal",
		"ruby", "sapphire", "emerald", "firered", "leafgreen",
		"diamond", "pearl", "platinum", "heartgold", "soulsilver", "black", "white", "black-2", "white-2",
	}
	
	for _, flavorText := range species.FlavorTextEntries {
		// Only include english flavor texts, whose versions are not in 'omissions'
		if (flavorText.Language.Name == "en" && !slices.Contains(omissions, flavorText.Version.Name)) {
			flavorTexts = append(flavorTexts, flavorText)
		}
	}

	flavorTexts = lo.UniqBy(flavorTexts, func(e FlavorText) string {
		return e.FlavorText
	})

	data := PkmnData{
		Pokemon: pkmn,
		PaddedID: paddedID, 
		EnglishGenus: engGenus.Genus, 
		FlavorTexts: flavorTexts,
	}

	serverSassComp(false)

	parseTemp("pkmn.tmpl").Execute(w, data)

}