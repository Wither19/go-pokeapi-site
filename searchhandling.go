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

func searchExactNumber(w http.ResponseWriter, r *http.Request, searchTerm string) {
	http.Redirect(w, r, fmt.Sprintf("/pkmn/%v", searchTerm), http.StatusFound)
}

func searchExactName(w http.ResponseWriter, r *http.Request, searchTerm string) {
	searchedPkmnNumber := slices.Index(pkmnNames, searchTerm) + 1
	http.Redirect(w, r, fmt.Sprintf("/pkmn/%d", searchedPkmnNumber), http.StatusFound)
}

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

func searchSubstring(searchTerm string) []NatlDexEntry {
	d := lo.Filter(natlDexEntries, func(item NatlDexEntry, _ int) bool {
		return strings.Contains(strings.ToLower(item.PokemonSpecies.Name), searchTerm)
	})

	return d
}

func pkmnSearchHandle(w http.ResponseWriter, r *http.Request, searchTerm string) []NatlDexEntry {
	searchTerm = strings.ToLower(searchTerm)
	var d []NatlDexEntry

	// 1. Exact Number
	// 2. Exact Name
	// 3. Range between two Dex #s, delimited by a dash
	// 4. Filter by region
	// 5. Regular Substring Match
	if _, err := strconv.ParseInt(searchTerm, 0, 0); err == nil {
		searchExactNumber(w, r, searchTerm)

	} else if slices.Contains(pkmnNames, searchTerm) {
		searchExactName(w, r, searchTerm)

	} else if strings.Contains(searchTerm, "-") {
		d = searchRange(searchTerm)

	} else if slices.Contains(regionKeywords, searchTerm) {
		d = searchRegion(searchTerm)

	} else {
		d = searchSubstring(searchTerm)
	}

	return d
}
