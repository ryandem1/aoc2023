// --- Day 8: Haunted Wasteland ---
// https://adventofcode.com/2023/day/8

package solutions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Day8Part1(inputPath string) (string, error) {
	// Starting at AAA, follow the left/right instructions.
	// How many steps are required to reach ZZZ?
	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	var instructions string
	network := make(map[string][]string) // network graph as adjacency list

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
	}
	err = file.Close()
	if err != nil {
		return "", err
	}

	fmt.Println(network)
	err = scanner.Err()
	if err != nil {
		return "", err
	}
	return "", nil
}

func Day8Part2(inputPath string) (string, error) {
	return "", nil
}
