"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.pkmnSearch = pkmnSearch;
function pkmnSearch(searchTerm = "") {
    if (!searchTerm) {
        const searchBar = document.getElementById("pkmn-search-bar");
        searchTerm = searchBar.value;
    }
    window.location.href = `/search/${searchTerm}`;
}
document.getElementById("pkmn-search-btn").addEventListener("click", () => {
    pkmnSearch();
});
document.getElementById("pkmn-search-bar").addEventListener("keydown", (e) => {
    if (e.key == "Enter") {
        pkmnSearch();
    }
});
