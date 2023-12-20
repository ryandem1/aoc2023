// --- Day 8: Haunted Wasteland ---
// https://adventofcode.com/2023/day/8

package solutions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Day8Part1(inputPath string) (string, error) {
	// Starting at AAA, follow the left/right instructions.
	// How many steps are required to reach ZZZ?

	stepsRequired := 0
	var instructions string
	var currNode string
	targetNode := "ZZZ"
	network := make(map[string][]string) // network graph as adjacency list
	instToIdx := map[uint8]int{'L': 0, 'R': 1}

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if instructions == "" {
			instructions = line
			continue
		}

		if line == "" {
			continue
		}

		var source string
		var sink1 string
		var sink2 string

		numSuccess, err := fmt.Sscanf(line, "%3s = (%3s, %3s)", &source, &sink1, &sink2)
		if err != nil {
			return "", err
		}

		if numSuccess != 3 {
			return "", errors.New("did not fully parse network")
		}

		network[source] = []string{sink1, sink2}

		if currNode == "" {
			currNode = source // currNode has to initialize at the first source node
		}
	}
	err = scanner.Err()
	if err != nil {
		return "", err
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	i := 0
	for currNode != targetNode {
		stepsRequired++
		inst := instructions[i]
		idx := instToIdx[inst]

		currNode = network[currNode][idx]

		i++
		if i >= len(instructions) {
			i = 0
		}
	}
	return strconv.Itoa(stepsRequired), nil
}

func Day8Part2(inputPath string) (string, error) {
	return "", nil
}
