function pkmnSearch() {
	const searchTerm = document.getElementById("pkmn-search-bar").value;
	window.location.href = `/search/${searchTerm}`;
}


document.getElementById("pkmn-search-btn").addEventListener("click", pkmnSearch)
document.getElementById("pkmn-search-bar").addEventListener("submit", pkmnSearch);