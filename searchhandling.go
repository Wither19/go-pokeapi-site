package main

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/samber/lo"
)

// Redirects to the Pokemon's page if it's an exact number
func searchExactNumber(w http.ResponseWriter, r *http.Request, searchTerm string) {
	http.Redirect(w, r, fmt.Sprintf("/pkmn/%v", searchTerm), http.StatusFound)
}

// Redirects based on Pokemon's name
func searchExactName(w http.ResponseWriter, r *http.Request, searchTerm string) {
	namedPkmnIndex := slices.Index(pkmnNames, searchTerm) + 1

	http.Redirect(w, r, fmt.Sprintf("/pkmn/%d", namedPkmnIndex), http.StatusFound)
}

// Pulls up list of Pokemon in a range, delimited by a dash
func searchRange(searchTerm string) []NationalDexEntry {
	numRange := strings.Split(searchTerm, "-")

	start := parseInt(numRange[0])
	end := parseInt(numRange[1])

	d := lo.Filter(natlDexEntries, func(item NationalDexEntry, _ int) bool {
		return item.EntryNumber >= start && item.EntryNumber <= end
	})

	return d
}

// Pulls up list of Pokemon in a region
func searchRegion(searchTerm string) []NationalDexEntry {
	d := lo.Filter(natlDexEntries, func(item NationalDexEntry, _ int) bool {
		var regionConditional bool
		searchIndex := slices.Index(regionKeywords, searchTerm)

		startValue := regionStarts[searchIndex]
		endValue := regionStarts[searchIndex+1]

		regionConditional = item.EntryNumber >= startValue && item.EntryNumber < endValue

		return regionConditional
	})

	return d
}

// Generic substring search
func searchSubstring(searchTerm string) []NationalDexEntry {
	d := lo.Filter(natlDexEntries, func(item NationalDexEntry, _ int) bool {
		return strings.Contains(strings.ToLower(item.PokemonSpecies.Name), searchTerm)
	})

	return d
}
