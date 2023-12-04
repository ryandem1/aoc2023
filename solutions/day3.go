// Day 3: Gear Ratios
// https://adventofcode.com/2023/day/3

package solutions

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Day3Part1(inputPath string) (string, error) {
	// What is the sum of all the part numbers in the engine schematic?
	sumOfPartNumbers := 0

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var engineSchematic []rune
	lineLen := 0
	seenFirstNewline := false

	r := bufio.NewReader(file)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		} else {
			if c == '\n' && !seenFirstNewline {
				seenFirstNewline = true
				lineLen += 1
			}
			if !seenFirstNewline {
				lineLen += 1
			}
			engineSchematic = append(engineSchematic, c)
		}
	}

	isSymbol := func(c rune) bool {
		return !unicode.IsDigit(c) && c != '.' && c != '\n'
	}
	var currentNumber strings.Builder

	for i, c := range engineSchematic {
		if unicode.IsDigit(c) {
			currentNumber.WriteRune(c)
		} else if currentNumber.Len() > 0 {
			isPart := false
			partNum, err := strconv.Atoi(currentNumber.String())
			if err != nil {
				return "", err
			}

			iNumStart := i - currentNumber.Len()
			iNumEnd := i - 1
			iNumTopLeft := iNumStart - 1 - lineLen
			iNumTopRight := iNumEnd + 1 - lineLen
			iNumBotLeft := iNumStart - 1 + lineLen
			iNumBotRight := iNumEnd + 1 + lineLen

			// Check left for symbol
			if iNumStart-1 > 0 && isSymbol(engineSchematic[iNumStart-1]) {
				isPart = true
			}

			// Check right for symbol
			if isSymbol(engineSchematic[iNumEnd+1]) {
				isPart = true
			}

			// Check top for symbol
			if iNumTopLeft > 0 && isSymbol(engineSchematic[iNumTopLeft]) {
				isPart = true
			}

			if iNumTopRight > 0 && isSymbol(engineSchematic[iNumTopRight]) {
				isPart = true
			}

			if iNumTopLeft+1 > 0 {
				for j := iNumTopLeft + 1; j < iNumTopRight; j++ {
					if isSymbol(engineSchematic[j]) {
						isPart = true
					}
				}
			}

			// Check bottom for symbol
			if iNumBotLeft < len(engineSchematic) && isSymbol(engineSchematic[iNumBotLeft]) {
				isPart = true
			}

			if iNumBotRight < len(engineSchematic) && isSymbol(engineSchematic[iNumBotRight]) {
				isPart = true
			}

			if iNumBotLeft+1 < len(engineSchematic) {
				for j := iNumBotLeft + 1; j < iNumBotRight; j++ {
					if isSymbol(engineSchematic[j]) {
						isPart = true
					}
				}
			}

			if isPart {
				sumOfPartNumbers += partNum
			}
			currentNumber.Reset()
		}
	}

	return strconv.Itoa(sumOfPartNumbers), nil
}

func Day3Part2(inputPath string) (string, error) {
	// What is the sum of all the gear ratios in your engine schematic?
	sumOfGearRatios := 0

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var engineSchematic []rune
	lineLen := 0
	seenFirstNewline := false

	r := bufio.NewReader(file)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		} else {
			if c == '\n' && !seenFirstNewline {
				seenFirstNewline = true
				lineLen += 1
			}
			if !seenFirstNewline {
				lineLen += 1
			}
			engineSchematic = append(engineSchematic, c)
		}
	}

	isGear := func(c rune) bool {
		return c == '*'
	}
	var currentNumber strings.Builder
	gearValues := make(map[int][]int)
	appendGearValue := func(gear int, val int) {
		_, ok := gearValues[gear]
		if !ok {
			gearValues[gear] = []int{val}
		} else {
			gearValues[gear] = append(gearValues[gear], val)
		}
	}

	for i, c := range engineSchematic {
		if unicode.IsDigit(c) {
			currentNumber.WriteRune(c)
		} else if currentNumber.Len() > 0 {
			partNum, err := strconv.Atoi(currentNumber.String())
			if err != nil {
				return "", err
			}

			iNumStart := i - currentNumber.Len()
			iNumEnd := i - 1
			iNumTopLeft := iNumStart - 1 - lineLen
			iNumTopRight := iNumEnd + 1 - lineLen
			iNumBotLeft := iNumStart - 1 + lineLen
			iNumBotRight := iNumEnd + 1 + lineLen

			// Check left for symbol
			if iNumStart-1 > 0 && isGear(engineSchematic[iNumStart-1]) {
				appendGearValue(iNumStart-1, partNum)
			}

			// Check right for symbol
			if isGear(engineSchematic[iNumEnd+1]) {
				appendGearValue(iNumEnd+1, partNum)
			}

			// Check top for symbol
			if iNumTopLeft > 0 && isGear(engineSchematic[iNumTopLeft]) {
				appendGearValue(iNumTopLeft, partNum)
			}

			if iNumTopRight > 0 && isGear(engineSchematic[iNumTopRight]) {
				appendGearValue(iNumTopRight, partNum)
			}

			if iNumTopLeft+1 > 0 {
				for j := iNumTopLeft + 1; j < iNumTopRight; j++ {
					if isGear(engineSchematic[j]) {
						appendGearValue(j, partNum)
					}
				}
			}

			// Check bottom for symbol
			if iNumBotLeft < len(engineSchematic) && isGear(engineSchematic[iNumBotLeft]) {
				appendGearValue(iNumBotLeft, partNum)
			}

			if iNumBotRight < len(engineSchematic) && isGear(engineSchematic[iNumBotRight]) {
				appendGearValue(iNumBotRight, partNum)
			}

			if iNumBotLeft+1 < len(engineSchematic) {
				for j := iNumBotLeft + 1; j < iNumBotRight; j++ {
					if isGear(engineSchematic[j]) {
						appendGearValue(j, partNum)
					}
				}
			}

			currentNumber.Reset()
		}
	}

	for _, values := range gearValues {
		if len(values) == 2 {
			sumOfGearRatios += values[0] * values[1]
		}
	}

	return strconv.Itoa(sumOfGearRatios), nil
}
