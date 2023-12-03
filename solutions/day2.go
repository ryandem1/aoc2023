// Day 2: Cube Conundrum
// https://adventofcode.com/2023/day/2

package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Day2Part1(inputPath string) (string, error) {
	// Determine which games would have been possible if the bag had been
	// loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes.
	// What is the sum of the IDs of those games?
	cubeCapacityByColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sumPossibleGameIDs := 0

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		colonSeen := false
		gamePossible := true
		var gameID strings.Builder
		var quantity strings.Builder
		var color strings.Builder

		evalCubes := func() error {
			quantityValue, err := strconv.Atoi(quantity.String())
			if err != nil {
				return err
			}

			capacity, ok := cubeCapacityByColor[color.String()]
			if !ok {
				return fmt.Errorf("invalid color: %s", color.String())
			}

			if quantityValue > capacity {
				gamePossible = false
			}
			quantity.Reset()
			color.Reset()

			return nil
		}

		for _, char := range line {

			if char == ':' {
				colonSeen = true
			}

			if unicode.IsDigit(char) && !colonSeen {
				gameID.WriteRune(char)
			}

			if unicode.IsDigit(char) && colonSeen {
				quantity.WriteRune(char)
			}

			if unicode.IsLetter(char) && colonSeen {
				color.WriteRune(char)
			}

			if char == ',' || char == ';' {
				err = evalCubes()
				if err != nil {
					return "", err
				}
			}
		}

		err = evalCubes()
		if err != nil {
			return "", err
		}

		gameIDValue, err := strconv.Atoi(gameID.String())
		if err != nil {
			return "", err
		}

		if gamePossible {
			sumPossibleGameIDs += gameIDValue
		}

		gameID.Reset()
	}

	if err = scanner.Err(); err != nil {
		return "", nil
	}

	return strconv.Itoa(sumPossibleGameIDs), nil
}

func Day2Part2(inputPath string) (string, error) {
	// For each game, find the minimum set of cubes that must have been present.
	// What is the sum of the power of these sets?
	sumOfSquaresOfGames := 0

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		minCubesNeeded := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		colonSeen := false
		var quantity strings.Builder
		var color strings.Builder

		// Helpful closure to evaluate the current quantity/color
		updateMinCubesNeeded := func() error {
			quantityValue, err := strconv.Atoi(quantity.String())
			if err != nil {
				return err
			}

			cubeCapByColor := minCubesNeeded[color.String()]
			if quantityValue > cubeCapByColor {
				minCubesNeeded[color.String()] = quantityValue
			}
			quantity.Reset()
			color.Reset()

			return nil
		}

		for _, char := range line {

			if char == ':' {
				colonSeen = true
			}

			if unicode.IsDigit(char) && colonSeen {
				quantity.WriteRune(char)
			}

			if unicode.IsLetter(char) && colonSeen {
				color.WriteRune(char)
			}

			if char == ',' || char == ';' {
				err = updateMinCubesNeeded()
				if err != nil {
					return "", err
				}
			}
		}

		err = updateMinCubesNeeded()
		if err != nil {
			return "", err
		}

		sumOfSquaresOfGames += minCubesNeeded["red"] * minCubesNeeded["blue"] * minCubesNeeded["green"]
	}

	if err = scanner.Err(); err != nil {
		return "", nil
	}

	return strconv.Itoa(sumOfSquaresOfGames), nil
}
