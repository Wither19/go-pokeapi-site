package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

// Redirects to the Pokemon's page if it's an exact number
func searchExactNumber(w http.ResponseWriter, r *http.Request, searchTerm string) {
	http.Redirect(w, r, fmt.Sprintf("/pkmn/%v", searchTerm), http.StatusFound)
}

// Redirects based on Pokemon's name
func searchExactName(w http.ResponseWriter, r *http.Request, searchTerm string) {
	http.Redirect(w, r, fmt.Sprintf("/pkmn/%d", slices.Index(pkmnNames, searchTerm)+1), http.StatusFound)
}

// Pulls up list of Pokemon in a range, delimited by a dash
func searchRange(searchTerm string) []NatlDexEntry {
	numRange := strings.Split(searchTerm, "-")

	start, startParseErr := strconv.ParseInt(numRange[0], 0, 0)
	if startParseErr != nil {
		log.Fatalln("Failed to parse start number in search range")
	}
	end, endParseErr := strconv.ParseInt(numRange[1], 0, 0)
	if endParseErr != nil {
		log.Fatalln("Failed to parse end number in search range")
	}

	d := lo.Filter(natlDexEntries, func(item NatlDexEntry, _ int) bool {
		return item.EntryNumber >= int(start) && item.EntryNumber <= int(end)
	})

	return d
}

// Pulls up list of Pokemon in a region
func searchRegion(searchTerm string) []NatlDexEntry {
	d := lo.Filter(natlDexEntries, func(item NatlDexEntry, _ int) bool {
		regionConditional := false
		startValue := regionStarts[slices.Index(regionKeywords, searchTerm)]
		if searchTerm == "paldea" {
			regionConditional = item.EntryNumber >= startValue && item.EntryNumber <= 1025
		} else {
			endValue := regionStarts[slices.Index(regionKeywords, searchTerm)+1]
			regionConditional = item.EntryNumber >= startValue && item.EntryNumber < endValue
		}

		return regionConditional
	})

	return d
}

// Generic substring search
func searchSubstring(searchTerm string) []NatlDexEntry {
	d := lo.Filter(natlDexEntries, func(item NatlDexEntry, _ int) bool {
		return strings.Contains(strings.ToLower(item.PokemonSpecies.Name), searchTerm)
	})

	return d
}

func pkmnSearchHandle(w http.ResponseWriter, r *http.Request, searchTerm string) []NatlDexEntry {
	searchTerm = strings.ToLower(searchTerm)
	var d []NatlDexEntry

	if _, err := strconv.ParseInt(searchTerm, 0, 0); err == nil {
		searchExactNumber(w, r, searchTerm)
	} else if slices.Contains(pkmnNames, searchTerm) {
		searchExactName(w, r, searchTerm)
	} else if strings.Contains(searchTerm, "-") {
		d = searchRange(searchTerm)
	} else if slices.Contains(regionKeywords, searchTerm) {
		d = searchRegion(searchTerm)
	} else if searchTerm == "random" {
		searchExactNumber(w, r, fmt.Sprint(randomNumber(1, 1026)))
	} else {
		d = searchSubstring(searchTerm)
	}

	return d
}
