.PHONY: frontend_build build dev preview

sass_dir := frontend/scss/App.scss
compiled_css_dir := static/css/style.css

frontend_build:
	tsc
	sass $(sass_dir) $(compiled_css_dir)

build: frontend_build
	go build

dev: frontend_build
	go run .

preview: build
	./jv-pokeapi