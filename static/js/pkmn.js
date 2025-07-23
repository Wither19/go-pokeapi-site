const regArt = document.querySelector("#regular-art");
const shinyArt = document.querySelector("#shiny-art");

function togglePkmnArt() {
	regArt.classList.toggle("hidden");
	shinyArt.classList.toggle("hidden");
}

regArt.addEventListener("click", togglePkmnArt);
shinyArt.addEventListener("click", togglePkmnArt);

function setFlavor(game) {
	const nonGameElems = document.querySelectorAll(`.dex-entries > div:not(#${game})`);
	const gameElem = document.querySelector(`div#${game}`);

	nonGameElems.forEach((elem) => elem.classList.add("hidden"));
	gameElem.classList.remove("hidden");
}