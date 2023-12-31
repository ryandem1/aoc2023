package main

import (
	"fmt"
	"github.com/ryandem1/aoc2023/solutions"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: main.go <day> <1|2> <example|full>")
	}

	day := os.Args[1]
	part := os.Args[2]
	dataset := os.Args[3]

	log.Printf("Running with dataset: %s", dataset)

	inputPath := fmt.Sprintf("data/day%s%s.txt", day, dataset)

	funcName := fmt.Sprintf("Day%sPart%s", day, part)

	dayFunc := map[string]func(args ...string) (string, error){
		"Day1Part1": solutions.Day1Part1,
		"Day1Part2": solutions.Day1Part2,
		"Day2Part1": solutions.Day2Part1,
		"Day2Part2": solutions.Day2Part2,
		"Day3Part1": solutions.Day3Part1,
		"Day3Part2": solutions.Day3Part2,
		"Day4Part1": solutions.Day4Part1,
		"Day4Part2": solutions.Day4Part2,
		"Day5Part1": solutions.Day5Part1,
		"Day5Part2": solutions.Day5Part2,
		"Day6Part1": solutions.Day6Part1,
		"Day6Part2": solutions.Day6Part2,
		"Day7Part1": solutions.Day7Part1,
		"Day7Part2": solutions.Day7Part2,
		"Day8Part1": solutions.Day8Part1,
		"Day8Part2": solutions.Day8Part2,
		"Day9Part1": solutions.Day9Part1,
		"Day9Part2": solutions.Day9Part2,
	}[funcName]

	if dayFunc == nil {
		log.Fatalf("Unrecognized day/part: %s", funcName)
	}

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	startAlloc := memStats.Alloc

	startTime := time.Now()
	solution, err := dayFunc(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	elapsedTime := time.Since(startTime)

	runtime.ReadMemStats(&memStats)
	endAlloc := memStats.Alloc

	log.Printf("Day %s Part %s Solution: %s", day, part, solution)
	log.Printf("Execution Time: %s", elapsedTime)
	log.Printf("Memory Used: %d KB", (endAlloc-startAlloc)/1024)

	// Delete temp file if created
	if _, err := os.Stat("temp.txt"); err == nil {
		err = os.Remove("temp.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
}
