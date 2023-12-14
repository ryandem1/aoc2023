// Day 7: Camel Cards
// https://adventofcode.com/2023/day/7

package solutions

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Day7Part1(inputPath string) (string, error) {
	// Find the rank of every hand in your set.
	// What are the total winnings?
	totalWinnings := 0
	handsByScore := make(map[int][]string)
	cardScores := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	getHandScore := func(hand string) int {
		cardCounts := make(map[rune]int)
		for _, card := range hand {
			cardCounts[card] += 1
		}
		highestCount := 0
		for _, count := range cardCounts {
			if count > highestCount {
				highestCount = count
			}
		}

		// Five-of-a-kind
		if highestCount == 5 {
			return 7
		}

		// Four-of-a-kind
		if highestCount == 4 {
			return 6
		}

		// Full-house
		if highestCount == 3 && len(cardCounts) == 2 {
			return 5
		}

		// Three-of-a-kind
		if highestCount == 3 {
			return 4
		}

		// Two-pair
		if highestCount == 2 && len(cardCounts) == 3 {
			return 3
		}

		// One-pair
		if highestCount == 2 {
			return 2
		}

		return 1
	}

	breakTies := func(hands []string) (sortedHands []string) {
		scoresToCard := make(map[int]rune)

		for card, score := range cardScores {
			scoresToCard[score] = card
		}

		handValues := make([][]int, len(hands))

		for i, hand := range hands {
			for _, card := range hand {
				handValues[i] = append(handValues[i], cardScores[card])
			}
		}

		vIndex := 0
		var sortHands func(i, j int) bool
		sortHands = func(i, j int) bool {
			h1, h2 := handValues[i], handValues[j]
			valueI := h1[vIndex]
			valueJ := h2[vIndex]

			if valueI == valueJ {
				vIndex++
				return sortHands(i, j)
			}

			s := valueI > valueJ
			vIndex = 0
			return s
		}

		sort.SliceStable(handValues, sortHands)

		for _, handValue := range handValues {
			var hand []rune
			for _, score := range handValue {
				card := scoresToCard[score]
				hand = append(hand, card)
			}

			sortedHands = append(sortedHands, string(hand))
		}
		slices.Reverse(sortedHands)
		return sortedHands
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	handBids := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		hand, bidStr := splitLine[0], splitLine[1]
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			return "", err
		}

		handScore := getHandScore(hand)
		handBids[hand] = bid
		handsByScore[handScore] = append(handsByScore[handScore], hand)
	}

	var handsSortedByRank []string
	for score := 1; score < 8; score++ {
		hands := handsByScore[score]
		if len(hands) > 0 {
			sortedHands := breakTies(hands)
			for _, hand := range sortedHands {
				handsSortedByRank = append(handsSortedByRank, hand)
			}
		}
	}

	for i, hand := range handsSortedByRank {
		rank := i + 1
		bid := handBids[hand]
		totalWinnings += bid * rank
	}

	err = scanner.Err()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(totalWinnings), nil
}

func Day7Part2(inputPath string) (string, error) {
	// Using the new joker rule, find the rank of every hand in your set.
	// What are the new total winnings?
	totalWinnings := 0
	handsByScore := make(map[int][]string)
	cardScores := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

	getHandScore := func(hand string) int {
		cardCounts := make(map[rune]int)
		numJokers := 0
		for _, card := range hand {
			if card == 'J' {
				numJokers++
				continue
			}
			cardCounts[card] += 1
		}
		highestCount := 0
		highestCountCard := 'J'

		for card, count := range cardCounts {
			if count > highestCount {
				highestCount = count
				highestCountCard = card
			} else if count == highestCount {
				if cardScores[card] > cardScores[highestCountCard] {
					highestCountCard = card
				}
			}
		}
		hand = strings.ReplaceAll(hand, "J", string(highestCountCard))
		highestCount += numJokers
		cardCounts[highestCountCard] += numJokers

		// Five-of-a-kind
		if highestCount == 5 {
			return 7
		}

		// Four-of-a-kind
		if highestCount == 4 {
			return 6
		}

		// Full-house
		if highestCount == 3 && len(cardCounts) == 2 {
			return 5
		}

		// Three-of-a-kind
		if highestCount == 3 {
			return 4
		}

		// Two-pair
		if highestCount == 2 && len(cardCounts) == 3 {
			return 3
		}

		// One-pair
		if highestCount == 2 {
			return 2
		}

		return 1
	}

	breakTies := func(hands []string) (sortedHands []string) {
		scoresToCard := make(map[int]rune)

		for card, score := range cardScores {
			scoresToCard[score] = card
		}

		handValues := make([][]int, len(hands))

		for i, hand := range hands {
			for _, card := range hand {
				handValues[i] = append(handValues[i], cardScores[card])
			}
		}

		vIndex := 0
		var sortHands func(i, j int) bool
		sortHands = func(i, j int) bool {
			h1, h2 := handValues[i], handValues[j]
			valueI := h1[vIndex]
			valueJ := h2[vIndex]

			if valueI == valueJ {
				vIndex++
				return sortHands(i, j)
			}

			s := valueI > valueJ
			vIndex = 0
			return s
		}

		sort.SliceStable(handValues, sortHands)

		for _, handValue := range handValues {
			var hand []rune
			for _, score := range handValue {
				card := scoresToCard[score]
				hand = append(hand, card)
			}

			sortedHands = append(sortedHands, string(hand))
		}
		slices.Reverse(sortedHands)
		return sortedHands
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	handBids := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		hand, bidStr := splitLine[0], splitLine[1]
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			return "", err
		}

		handScore := getHandScore(hand)
		handBids[hand] = bid
		handsByScore[handScore] = append(handsByScore[handScore], hand)
	}

	var handsSortedByRank []string
	for score := 1; score < 8; score++ {
		hands := handsByScore[score]
		if len(hands) > 0 {
			sortedHands := breakTies(hands)
			for _, hand := range sortedHands {
				handsSortedByRank = append(handsSortedByRank, hand)
			}
		}
	}

	for i, hand := range handsSortedByRank {
		rank := i + 1
		bid := handBids[hand]
		totalWinnings += bid * rank
	}

	err = scanner.Err()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(totalWinnings), nil
}
