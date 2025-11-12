const searchBar = document.querySelector("#pkmn-search-bar")! as HTMLInputElement
const searchBtn = document.querySelector("#pkmn-search-btn")! as HTMLButtonElement

export function pkmnSearch(searchTerm = searchBar.value) {
	window.location.href = `/search/${searchTerm}`
}

searchBtn.addEventListener("click", () => {
	pkmnSearch()
})

searchBar.addEventListener("keydown", (e) => {
	if (e.key == "Enter") {
		pkmnSearch()
	}
})

