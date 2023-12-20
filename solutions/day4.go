// Day 4: Scratchcards
// https://adventofcode.com/2023/day/4
// I am not 100% proud of this lol

package solutions

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day4Part1(args ...string) (string, error) {
	// Take a seat in the large pile of colorful cards.
	// How many points are they worth in total?
	inputPath := args[0]
	totalPoints := 0

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	isExampleFile := strings.Contains(inputPath, "example")

	var numbersOnCard int
	var numberOfWinningNumbers int

	if isExampleFile {
		numbersOnCard = 8
		numberOfWinningNumbers = 5
	} else {
		numbersOnCard = 25
		numberOfWinningNumbers = 10
	}

	cardID := 0
	winningNumbers := make([]int, numberOfWinningNumbers)
	cardNumbers := make([]int, numbersOnCard)

	// Infinite loop to read and parse each line
	for {
		if isExampleFile {
			_, err = fmt.Fscanf(file, "Card %d: %d %d %d %d %d | %d %d %d %d %d %d %d %d\n",
				&cardID,
				&winningNumbers[0], &winningNumbers[1], &winningNumbers[2], &winningNumbers[3], &winningNumbers[4],
				&cardNumbers[0], &cardNumbers[1], &cardNumbers[2], &cardNumbers[3], &cardNumbers[4], &cardNumbers[5],
				&cardNumbers[6], &cardNumbers[7])
		} else {
			_, err = fmt.Fscanf(file, "Card %d: %d %d %d %d %d %d %d %d %d %d | %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d\n",
				&cardID,
				&winningNumbers[0], &winningNumbers[1], &winningNumbers[2], &winningNumbers[3], &winningNumbers[4],
				&winningNumbers[5], &winningNumbers[6], &winningNumbers[7], &winningNumbers[8], &winningNumbers[9],
				&cardNumbers[0], &cardNumbers[1], &cardNumbers[2], &cardNumbers[3], &cardNumbers[4],
				&cardNumbers[5], &cardNumbers[6], &cardNumbers[7], &cardNumbers[8], &cardNumbers[9],
				&cardNumbers[10], &cardNumbers[11], &cardNumbers[12], &cardNumbers[13], &cardNumbers[14],
				&cardNumbers[15], &cardNumbers[16], &cardNumbers[17], &cardNumbers[18], &cardNumbers[19],
				&cardNumbers[20], &cardNumbers[21], &cardNumbers[22], &cardNumbers[23], &cardNumbers[24])
		}

		intersection := make([]int, 0)

		uniqueElements := make(map[int]bool)
		for _, num := range winningNumbers {
			uniqueElements[num] = true
		}

		for _, num := range cardNumbers {
			if uniqueElements[num] {
				intersection = append(intersection, num)
			}
		}

		totalPoints += int(math.Pow(float64(2), float64(len(intersection)-1)))

		// Check for end of file
		if errors.Is(err, io.ErrUnexpectedEOF) {
			break
		} else if err != nil {
			return "", err
		}
	}
	return strconv.Itoa(totalPoints), nil
}

func Day4Part2(args ...string) (string, error) {
	// Including the original set of scratchcards,
	// How many total scratchcards do you end up with?
	inputPath := args[0]
	totalCards := 0

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	isExampleFile := strings.Contains(inputPath, "example")

	var numbersOnCard int
	var numberOfWinningNumbers int

	if isExampleFile {
		numbersOnCard = 8
		numberOfWinningNumbers = 5
	} else {
		numbersOnCard = 25
		numberOfWinningNumbers = 10
	}

	cardID := 0
	winningNumbers := make([]int, numberOfWinningNumbers)
	cardNumbers := make([]int, numbersOnCard)
	countOfCard := make(map[int]int)

	for {
		if isExampleFile {
			_, err = fmt.Fscanf(file, "Card %d: %d %d %d %d %d | %d %d %d %d %d %d %d %d\n",
				&cardID,
				&winningNumbers[0], &winningNumbers[1], &winningNumbers[2], &winningNumbers[3], &winningNumbers[4],
				&cardNumbers[0], &cardNumbers[1], &cardNumbers[2], &cardNumbers[3], &cardNumbers[4], &cardNumbers[5],
				&cardNumbers[6], &cardNumbers[7])
		} else {
			_, err = fmt.Fscanf(file, "Card %d: %d %d %d %d %d %d %d %d %d %d | %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d\n",
				&cardID,
				&winningNumbers[0], &winningNumbers[1], &winningNumbers[2], &winningNumbers[3], &winningNumbers[4],
				&winningNumbers[5], &winningNumbers[6], &winningNumbers[7], &winningNumbers[8], &winningNumbers[9],
				&cardNumbers[0], &cardNumbers[1], &cardNumbers[2], &cardNumbers[3], &cardNumbers[4],
				&cardNumbers[5], &cardNumbers[6], &cardNumbers[7], &cardNumbers[8], &cardNumbers[9],
				&cardNumbers[10], &cardNumbers[11], &cardNumbers[12], &cardNumbers[13], &cardNumbers[14],
				&cardNumbers[15], &cardNumbers[16], &cardNumbers[17], &cardNumbers[18], &cardNumbers[19],
				&cardNumbers[20], &cardNumbers[21], &cardNumbers[22], &cardNumbers[23], &cardNumbers[24])
		}

		if errors.Is(err, io.ErrUnexpectedEOF) {
			break
		} else if err != nil {
			return "", err
		}

		countOfCard[cardID] += 1
		intersection := make([]int, 0)

		uniqueElements := make(map[int]bool)
		for _, num := range winningNumbers {
			uniqueElements[num] = true
		}

		for _, num := range cardNumbers {
			if uniqueElements[num] {
				intersection = append(intersection, num)
			}
		}

		for i := 0; i < len(intersection); i++ {
			countOfCard[i+cardID+1] += countOfCard[cardID]
		}
	}

	for _, count := range countOfCard {
		totalCards += count
	}
	return strconv.Itoa(totalCards), nil
}
