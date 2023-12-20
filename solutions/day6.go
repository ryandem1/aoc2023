// Day 6: Wait For It
// https://adventofcode.com/2023/day/6

package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day6Part1(args ...string) (string, error) {
	// Determine the number of ways you could beat the record in each race.
	// What do you get if you multiply these numbers together?
	inputPath := args[0]
	multOfAllPossibleWays := 1
	var times []int
	var distances []int

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		for _, item := range lineSplit {
			cleanedItem := strings.Trim(item, " ")
			itemVal, err := strconv.Atoi(cleanedItem)
			if err != nil {
				continue
			}
			if lineNum == 0 {
				times = append(times, itemVal)
			} else if lineNum == 1 {
				distances = append(distances, itemVal)
			} else {
				continue
			}
		}

		lineNum++
	}

	for i := 0; i < len(times); i++ {
		totalTime := times[i]
		goalDist := distances[i]

		lBound, uBound := -1, -1
		for chargeTime := 0; chargeTime < totalTime; chargeTime++ {
			dist := chargeTime * (totalTime - chargeTime)
			if dist > goalDist {
				lBound = chargeTime - 1
				break
			}
		}

		for chargeTime := totalTime; chargeTime > 0; chargeTime-- {
			dist := chargeTime * (totalTime - chargeTime)
			if dist > goalDist {
				uBound = chargeTime
				break
			}
		}

		possibleWays := uBound - lBound
		multOfAllPossibleWays *= possibleWays
	}

	err = scanner.Err()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(multOfAllPossibleWays), nil
}

func Day6Part2(args ...string) (string, error) {
	// How many ways can you beat the record in this one much longer race?
	inputPath := args[0]
	fileLines := ""

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cleanedLine := ""
		line := scanner.Text()
		splitLine := strings.Split(line, ":")

		cleanedLine += splitLine[0] + ": "
		cleanedLine += strings.ReplaceAll(splitLine[1], " ", "")
		fileLines += cleanedLine + "\n"
	}

	err = scanner.Err()
	if err != nil {
		return "", nil
	}

	err = os.WriteFile("temp.txt", []byte(fileLines), 0644)
	if err != nil {
		return "", err
	}

	return Day6Part1("temp.txt")
}
