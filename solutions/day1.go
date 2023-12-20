// Day 1: Trebuchet
// https://adventofcode.com/2023/day/1

package solutions

import (
	"os"
	"strconv"
	"strings"
)

func Day1Part1(args ...string) (string, error) {
	inputPath := args[0]

	const (
		byte0       = 48
		byte9       = 57
		byteNewline = 10
	)

	input, err := os.ReadFile(inputPath)
	if err != nil {
		return "", err
	}

	var firstDigit byte
	var lastDigit byte

	sumOfCalibrationValues := 0
	for _, charByte := range input {
		if charByte >= byte0 && charByte <= byte9 { // Digit range
			if firstDigit == 0 {
				firstDigit = charByte
				lastDigit = firstDigit
			} else {
				lastDigit = charByte
			}
		} else if charByte == byteNewline {
			var strBuilder strings.Builder

			strBuilder.WriteByte(firstDigit)
			strBuilder.WriteByte(lastDigit)

			calibrationValueString := strBuilder.String()
			calibrationValue, err := strconv.Atoi(calibrationValueString)
			if err != nil {
				return "", err
			}
			sumOfCalibrationValues += calibrationValue
			firstDigit = 0
			lastDigit = 0
		}
	}
	return strconv.Itoa(sumOfCalibrationValues), nil
}

func Day1Part2(args ...string) (string, error) {
	inputPath := args[0]

	content, err := os.ReadFile(inputPath)
	if err != nil {
		return "", err
	}

	text := string(content)

	numTexts := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, numText := range numTexts {
		var builder strings.Builder

		// Replace the number, but also need to preserve the first and last letter
		// of the line in case of overlap
		builder.WriteByte(numText[0])
		builder.WriteString(strconv.Itoa(i + 1))
		builder.WriteByte(numText[len(numText)-1])

		numWithChar := builder.String()
		text = strings.Replace(text, numText, numWithChar, -1)
	}

	err = os.WriteFile("temp.txt", []byte(text), 0644)
	if err != nil {
		return "", err
	}

	return Day1Part1("temp.txt")
}
