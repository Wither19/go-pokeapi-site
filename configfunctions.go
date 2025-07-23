package main

import (
	"log"
	"os"

	"github.com/stretchr/testify/assert/yaml"
)

func loadServerYAML() ServerConfig {
	var config ServerConfig

	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalln("Could not open YAML file:", err)
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalln("Could not unmarshal YAML file:", err)
	}

	return config
}
