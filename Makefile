.PHONY: sass_build build dev preview

sass_dir := scss/App.scss
compiled_css_dir := static/css/style.css

sass_build:
	@echo "Building SASS in $(sass_dir)"

	sass $(sass_dir) $(compiled_css_dir)

build: sass_build
	@echo "Building Go source"
	go build

dev: sass_build
	go run .

preview: build
	./jv-pokeapi