// Day 1: Trebuchet
// https://adventofcode.com/2023/day/1

package solutions

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	byte0       = 48
	byte9       = 57
	byteNewline = 10
)

func Day1Part1(inputPath string) string {
	// Consider your entire calibration document.
	// What is the sum of all the calibration values?
	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
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
				log.Fatal(err)
			}
			sumOfCalibrationValues += calibrationValue
			firstDigit = 0
			lastDigit = 0
		}
	}
	return strconv.Itoa(sumOfCalibrationValues)
}

func Day1Part2(inputPath string) string {
	return inputPath
}
