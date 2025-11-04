export function pkmnSearch(searchTerm: string = "") {
	if (!searchTerm) {
		const searchBar = document.getElementById("pkmn-search-bar")! as HTMLInputElement;
		searchTerm = searchBar.value;
	}
	window.location.href = `/search/${searchTerm}`;
}


document.getElementById("pkmn-search-btn")!.addEventListener("click", () => {
	pkmnSearch();
})
document.getElementById("pkmn-search-bar")!.addEventListener("keydown", (e) => {
	if (e.key == "Enter") {
		pkmnSearch();
	}
});
