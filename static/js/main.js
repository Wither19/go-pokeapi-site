"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const searchBar = document.querySelector("#pkmn-search-bar");
const searchBtn = document.querySelector("#pkmn-search-btn");
const pkmnSearch = (searchTerm = searchBar.value) => {
    window.location.href = `/search/${searchTerm}`;
};
searchBtn.addEventListener("click", () => {
    pkmnSearch();
});
searchBar.addEventListener("keydown", (e) => {
    if (e.key == "Enter") {
        pkmnSearch();
    }
});
