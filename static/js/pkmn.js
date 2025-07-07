const regArt = document.querySelector("#regular-art");
const shinyArt = document.querySelector("#shiny-art");

function togglePkmnArt() {
	regArt.classList.toggle("hidden");
	shinyArt.classList.toggle("hidden");
}

regArt.addEventListener("click", togglePkmnArt);
shinyArt.addEventListener("click", togglePkmnArt);

const flavorContainers = Array.from(document.querySelectorAll(".dex-entries > div"));
const flavorIds = flavorContainers.map((tag) => tag.id);

var currentFlavor = 0;

function cycleFlavor() {
	if (currentFlavor + 1 >= flavorIds.length) {
		currentFlavor = 0;
	} else {
		currentFlavor++;
	}

	for (let tagID of flavorIds) {
		const flavorElem = document.querySelector(`div#${tagID}`);

		if (flavorIds.indexOf(tagID) == currentFlavor) {
			flavorElem.classList.remove("hidden");
		} else {
			flavorElem.classList.add("hidden");
		}
	}
}

document.getElementById("flavor-cycle-btn").addEventListener("click", cycleFlavor);
