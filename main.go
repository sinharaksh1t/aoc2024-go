package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sinharaksh1t/aoc2024-go/day1"
	"github.com/sinharaksh1t/aoc2024-go/rootpath"
)

func main() {
	day1.Q1()
}

func init() {
	// Load .env file if it exists
	err := godotenv.Load(fmt.Sprintf("%s/.env", rootpath.Root))
	if err != nil {
		fmt.Println("No .env file found.")
	}
}
