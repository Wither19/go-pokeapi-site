"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const main_1 = require("./main");
const regArt = document.querySelector("#regular-art");
const shinyArt = document.querySelector("#shiny-art");
const togglePkmnArt = () => {
    regArt.classList.toggle("hidden");
    shinyArt.classList.toggle("hidden");
};
regArt.addEventListener("click", togglePkmnArt);
shinyArt.addEventListener("click", togglePkmnArt);
const setFlavor = (game) => {
    const nonGameElems = document.querySelectorAll(`.dex-entries > div:not(#${game})`);
    const gameElem = document.querySelector(`div#${game}`);
    nonGameElems.forEach((elem) => elem.classList.add("hidden"));
    gameElem.classList.remove("hidden");
};
document.querySelector("body").addEventListener("keydown", (e) => {
    generalKeyFuncs(e);
});
const generalKeyFuncs = (e) => {
    if (e.key == "r") {
        (0, main_1.pkmnSearch)("random");
    }
};
