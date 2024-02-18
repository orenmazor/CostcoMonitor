/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
)

func init() {
	LoadQueries()
}

func main() {
	b, _ := os.ReadFile("banner.txt")
	fmt.Printf(string(b))
}
