import { pkmnSearch } from "./main";

const regArt = document.querySelector("#regular-art")!;
const shinyArt = document.querySelector("#shiny-art")!;

const togglePkmnArt = () => {
	regArt.classList.toggle("hidden");
	shinyArt.classList.toggle("hidden");
};

regArt.addEventListener("click", togglePkmnArt);
shinyArt.addEventListener("click", togglePkmnArt);

const setFlavor = (game: string) => {
	const nonGameElems = document.querySelectorAll(`.dex-entries > div:not(#${game})`)!;
	const gameElem = document.querySelector(`div#${game}`)!;

	nonGameElems.forEach((elem) => elem.classList.add("hidden"));
	gameElem.classList.remove("hidden");
};

document.querySelector("body")!.addEventListener("keydown", (e) => {
	generalKeyFuncs(e);
});

const generalKeyFuncs = (e: KeyboardEvent) => {
	if (e.key == "r") {
			pkmnSearch("random");
	}
}