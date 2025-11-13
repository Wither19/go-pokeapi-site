"use strict";
const body = document.querySelector("body");
const searchBar = document.querySelector("#pkmn-search-bar");
const searchBtn = document.querySelector("#pkmn-search-btn");
const regArt = document.querySelector("#regular-art");
const shinyArt = document.querySelector("#shiny-art");
const pkmnSearch = (searchTerm = searchBar.value) => {
    window.location.href = `/search/${searchTerm}`;
};
const togglePkmnArt = () => {
    regArt.classList.toggle("hidden");
    shinyArt.classList.toggle("hidden");
};
const setFlavor = (game) => {
    const nonGameElems = document.querySelectorAll(`.dex-entries > div:not(#${game})`);
    const gameElem = document.querySelector(`div#${game}`);
    nonGameElems.forEach((elem) => {
        elem.classList.add("hidden");
        elem.classList.remove("hl");
    });
    gameElem.classList.remove("hidden");
    gameElem.classList.add("hl");
};
const generalKeyFuncs = (e) => {
    if (e.key == "r") {
        pkmnSearch("random");
    }
};
regArt.addEventListener("click", togglePkmnArt);
shinyArt.addEventListener("click", togglePkmnArt);
body.addEventListener("keydown", (e) => {
    generalKeyFuncs(e);
});
