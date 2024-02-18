/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log/slog"
	"os"
)

func init() {
	LoadQueries()
}

func main() {
	b, _ := os.ReadFile("banner.txt")
	fmt.Printf(string(b))

	for _, query := range SearchConfig.Queries {
		slog.Info("Checking query", "query", query.Query, "price limit", query.PriceLimit)

		GetCostcoResults(query)
	}
}
