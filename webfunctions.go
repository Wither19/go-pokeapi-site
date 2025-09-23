package main

import (
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/samber/lo"
)

// Returns a parsed template of the file provided by filename. If funcs is provided it gets added as a FuncMap to the template. Sprig's functions are loaded regardless.
func parseTemp(tempName string, fileName string, funcs template.FuncMap) *template.Template {
	var t *template.Template

	t = template.New(tempName)
	t = t.Funcs(sprig.FuncMap())

	if funcs != nil {
		t = t.Funcs(funcs)
	}

	return template.Must(t.ParseFiles(fileName))
}

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	parseTemp("main.html", "static/pages/main.html", nil).Execute(w, natlDexEntries)
}

func mainPagePkmnSearch(w http.ResponseWriter, r *http.Request) {
	searchTerm := strings.ToLower(r.PathValue("search"))
	var filteredDex []NationalDexEntry

	if _, err := strconv.ParseInt(searchTerm, 0, 0); err == nil {
		searchExactNumber(w, r, searchTerm)

	} else if slices.Contains(pkmnNames, searchTerm) {
		searchExactName(w, r, searchTerm)

	} else if searchTerm == "random" {
		searchExactNumber(w, r, fmt.Sprint(randomNumber(1, 1026)))

	} else if strings.Contains(searchTerm, "-") {
		filteredDex = searchRange(searchTerm)

	} else if slices.Contains(regionKeywords, searchTerm) {
		filteredDex = searchRegion(searchTerm)

	} else {
		filteredDex = searchSubstring(searchTerm)

	}

	parseTemp("main.html", "static/pages/main.html", nil).Execute(w, filteredDex)

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
		Config:       loadServerYAML(),
	}

	parseTemp("pkmn.html", "static/pages/pkmn.html", map[string]any{"leadingZeroes": leadingZeroes}).Execute(w, data)
}

func prevPkmnLoad(w http.ResponseWriter, r *http.Request) {
	currentPkmn := parseInt(strings.ReplaceAll(r.PathValue("id"), "/prev", "")) - 1

	prevPkmnUrl := fmt.Sprintf("/pkmn/%d", currentPkmn)

	http.Redirect(w, r, prevPkmnUrl, http.StatusFound)
}

func nextPkmnLoad(w http.ResponseWriter, r *http.Request) {
	currentPkmn := parseInt(strings.ReplaceAll(r.PathValue("id"), "/next", "")) + 1

	nextPkmnUrl := fmt.Sprintf("/pkmn/%d", currentPkmn)

	http.Redirect(w, r, nextPkmnUrl, http.StatusFound)
}

func pkmnRandomize(w http.ResponseWriter, r *http.Request) {
	randomID := randomNumber(1, 1026)
	randPkmnUrl := fmt.Sprintf("/pkmn/%v", randomID)

	http.Redirect(w, r, randPkmnUrl, http.StatusFound)
}
