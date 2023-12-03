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
				quantityValue, err := strconv.Atoi(quantity.String())
				if err != nil {
					return "", err
				}

				capacity, ok := cubeCapacityByColor[color.String()]
				if !ok {
					return "", fmt.Errorf("invalid color: %s", color.String())
				}

				if quantityValue > capacity {
					gamePossible = false
				}
				quantity.Reset()
				color.Reset()
			}
		}

		quantityValue, err := strconv.Atoi(quantity.String())
		if err != nil {
			return "", err
		}

		capacity, ok := cubeCapacityByColor[color.String()]
		if !ok {
			return "", fmt.Errorf("invalid color: %s", color.String())
		}

		if quantityValue > capacity {
			gamePossible = false
		}
		quantity.Reset()
		color.Reset()

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
	return "", nil
}
