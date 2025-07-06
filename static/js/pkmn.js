const regArt = document.querySelector("#regular-art");
const shinyArt = document.querySelector("#shiny-art");

function togglePkmnArt() {
	regArt.classList.toggle("hidden");
	shinyArt.classList.toggle("hidden");
}

regArt.addEventListener("click", togglePkmnArt);
shinyArt.addEventListener("click", togglePkmnArt);

const flavorTags = Array.from(document.querySelector(".dex-entries").children);
const flavorIds = flavorTags.map((tag) => {
	return tag.id;
});

var currentFlavor = 0;

function cycleFlavor() {
	if (currentFlavor + 1 >= flavorIds.length) {
		currentFlavor = 0;
	} else {
		currentFlavor++;
	}

	for (let tagID of flavorIds) {
		if (flavorIds.indexOf(tagID) == currentFlavor) {
			document.getElementById(tagID).classList.remove("hidden");
		} else {
			document.getElementById(tagID).classList.add("hidden");
		}
	}
}

for (let tagID of flavorIds) {
	document.getElementById("flavor-cycle-btn").addEventListener("click", cycleFlavor);
}
