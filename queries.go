package main

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

type Searches struct {
	Queries []Query `yaml:"queries"`
}

type Query struct {
	Query      string `yaml:"query"`
	PriceLimit int    `yaml:"price_limit"`
}

var SearchConfig Searches

func LoadQueries() {
	yamlFile, err := os.ReadFile("config/searches.yaml")
	if err != nil {
		panic(err)
	}

	// Parse the YAML file into a Config struct
	if err := yaml.Unmarshal(yamlFile, &SearchConfig); err != nil {
		panic(err)
	}

	slog.Info("Loaded query configuration", "queries found", len(SearchConfig.Queries))
}
