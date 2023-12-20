// --- Day 8: Haunted Wasteland ---
// https://adventofcode.com/2023/day/8

package solutions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func Day8Part1(args ...string) (string, error) {
	// Starting at AAA, follow the left/right instructions.
	// How many steps are required to reach ZZZ?

	inputPath := args[0]

	var currNode string
	if len(args) > 1 {
		currNode = args[1]
	}

	stepsRequired := 0
	var instructions string
	currNode = "AAA"
	targetNodes := []string{"ZZZ"}
	network := make(map[string]string) // network graph as adjacency list (sort of)

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

		network[source+"L"] = sink1
		network[source+"R"] = sink2
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
	for !slices.Contains(targetNodes, currNode) {
		stepsRequired++
		inst := instructions[i]

		currNode = network[currNode+string(inst)]

		i++
		if i >= len(instructions) {
			i = 0
		}
	}
	return strconv.Itoa(stepsRequired), nil
}

func Day8Part2(args ...string) (string, error) {
	// Simultaneously start on every node that ends with A.
	// How many steps does it take before you're only on nodes that end with Z?
	return args[0], nil
}
