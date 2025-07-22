package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"slices"
	"strconv"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/samber/lo"
)

// Returns a parsed template of the filename provided by n. If f is provided it gets added as a FuncMap to the template. Sprig's functions are loaded regardless.
func parseTemp(n string, f template.FuncMap) *template.Template {
	var t *template.Template

	t = template.New(n)
	t = t.Funcs(sprig.FuncMap())

	if f != nil {
		t = t.Funcs(f)
	}

	return template.Must(t.ParseFiles(n))
}

// Transpiles the SASS at sassSource, to newCss. Requires Dart Sass to be installed, and in your $PATH
func serverSassComp() {
	sassSource := "./static/scss/App.scss"
	newCss := "./static/css/style.css"

	sassBuild := exec.Command("sass", sassSource, newCss)

	if err := sassBuild.Run(); err != nil {
		log.Fatalln("Sass build error:", err)
	}
}

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	d := getNatlDex().PokemonEntries

	serverSassComp()

	parseTemp("main.html", nil).Execute(w, d)
}

func mainPagePkmnSearch(w http.ResponseWriter, r *http.Request) {
	d := getNatlDex().PokemonEntries
	searchTerm := r.PathValue("search")

	filteredDex := lo.Filter(d, func(item struct {
		EntryNumber    int "json:\"entry_number\""
		PokemonSpecies struct {
			Name string "json:\"name\""
			URL  string "json:\"url\""
		} "json:\"pokemon_species\""
	}, i int) bool {
		var c bool
		if _, err := strconv.ParseInt(searchTerm, 0, 0); err != nil {
			c = strings.Contains(item.PokemonSpecies.Name, searchTerm)
		} else {
			c = strings.Contains(fmt.Sprintf("%d", item.EntryNumber), searchTerm)
		}
		return c
	})

	serverSassComp()

	parseTemp("main.html", nil).Execute(w, filteredDex)

}

func pkmnSearchNotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusNotFound)
}

func pkmnLoad(w http.ResponseWriter, r *http.Request) {
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

	var flavorTexts []FlavorText

	omissions := []string{
		"red", "blue", "yellow", "gold", "silver", "crystal",
		"ruby", "sapphire", "emerald", "firered", "leafgreen",
		"diamond", "pearl", "platinum", "heartgold", "soulsilver", "black", "white", "black-2", "white-2",
	}

	for _, flavorText := range species.FlavorTextEntries {
		// Only include english flavor texts, whose versions are not in 'omissions'
		if flavorText.Language.Name == "en" && !slices.Contains(omissions, flavorText.Version.Name) {
			flavorTexts = append(flavorTexts, flavorText)
		}
	}

	flavorTexts = lo.UniqBy(flavorTexts, func(e FlavorText) string {
		return e.FlavorText
	})

	data := PkmnData{
		Pokemon:      pkmn,
		PaddedID:     paddedID,
		EnglishGenus: engGenus.Genus,
		FlavorTexts:  flavorTexts,
	}

	serverSassComp()

	parseTemp("pkmn.html", nil).Execute(w, data)
}

func prevPkmnLoad(w http.ResponseWriter, r *http.Request) {
	currentPkmn, _ := strconv.ParseInt(strings.ReplaceAll(r.PathValue("id"), "/prev", ""), 0, 0)
	currentPkmn -= 1

	prevPkmnUrl := fmt.Sprintf("/pkmn/%d", currentPkmn)

	http.Redirect(w, r, prevPkmnUrl, http.StatusFound)
}

func nextPkmnLoad(w http.ResponseWriter, r *http.Request) {
	currentPkmn, _ := strconv.ParseInt(strings.ReplaceAll(r.PathValue("id"), "/next", ""), 0, 0)
	currentPkmn += 1

	nextPkmnUrl := fmt.Sprintf("/pkmn/%d", currentPkmn)

	http.Redirect(w, r, nextPkmnUrl, http.StatusFound)
}

func pkmnRandomize(w http.ResponseWriter, r *http.Request) {
	randomID := randomNumber(1, 1026)
	randPkmnUrl := fmt.Sprintf("/pkmn/%v", randomID)

	http.Redirect(w, r, randPkmnUrl, http.StatusFound)
}
