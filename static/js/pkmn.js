const pkmnArt = document.querySelector("img.artwork");

const regularArt = document.querySelector("#regular-art");
const shinyArt = document.querySelector("#shiny-art");

function togglePkmnArt() {
	regularArt.classList.toggle("hidden");
	shinyArt.classList.toggle("hidden");
}

regularArt.addEventListener("click", togglePkmnArt);
shinyArt.addEventListener("click", togglePkmnArt);
