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

// Returns a parsed template of the filename provided by filename. If funcs is provided it gets added as a FuncMap to the template. Sprig's functions are loaded regardless.
func parseTemp(filename string, funcs template.FuncMap, parseSass bool) *template.Template {
	var t *template.Template

	t = template.New(filename)
	t = t.Funcs(sprig.FuncMap())

	if funcs != nil {
		t = t.Funcs(funcs)
	}

	if parseSass {
		serverSassComp()
	}

	return template.Must(t.ParseFiles(filename))
}

func serverSassComp() {
	sassBuild := exec.Command("sass", "./static/scss/App.scss", "./static/css/style.css")

	if err := sassBuild.Run(); err != nil {
		log.Fatalln("Sass build error:", err)
	}
}

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	parseTemp("main.html", nil, true).Execute(w, natlDexEntries)
}

func mainPagePkmnSearch(w http.ResponseWriter, r *http.Request) {

	filteredDex := pkmnSearchHandle(w, r, r.PathValue("search"))

	parseTemp("main.html", nil, true).Execute(w, filteredDex)

}

func pkmnSearchNotFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
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

	parseTemp("pkmn.html", nil, true).Execute(w, data)
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
