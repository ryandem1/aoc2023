package main

import (
	"fmt"
	"github.com/ryandem1/aoc2023/solutions"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: main.go <day> <1|2> <example|full>")
	}

	day := os.Args[1]
	part := os.Args[2]
	dataset := os.Args[3]

	switch dataset {
	case "example":
		log.Println("Running with example dataset")
	case "full":
		log.Println("Running with full dataset")
	default:
		log.Fatalf("Invalid dataset: %s", dataset)
	}

	inputPath := fmt.Sprintf("data/day%s%s.txt", day, dataset)

	funcName := fmt.Sprintf("Day%sPart%s", day, part)

	dayFunc := map[string]func(string) string{
		"Day1Part1": solutions.Day1Part1,
		"Day1Part2": solutions.Day1Part2,
	}[funcName]

	solution := dayFunc(inputPath)
	log.Printf("Day %s Part %s Solution: %s", day, part, solution)
}
