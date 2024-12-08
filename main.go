package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sinharaksh1t/aoc2024-go/day1"
	"github.com/sinharaksh1t/aoc2024-go/day2"
	"github.com/sinharaksh1t/aoc2024-go/day3"
	"github.com/sinharaksh1t/aoc2024-go/rootpath"
)

func main() {
	// TODO: Take input on which question to solve
	day1.Solve()
	day2.Solve()
	day3.Solve()
}

func init() {
	// Load .env file if it exists
	err := godotenv.Load(fmt.Sprintf("%s/.env", rootpath.Root))
	if err != nil {
		fmt.Println("No .env file found.")
	}
}
