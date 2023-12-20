// --- Day 8: Haunted Wasteland ---
// https://adventofcode.com/2023/day/8

package solutions

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strconv"
)

func Day8Part1(args ...string) (string, error) {
	inputPath := args[0]

	var currNode string
	if len(args) > 1 {
		currNode = args[1]
	} else {
		currNode = "AAA"
	}

	var targetNodes []string
	if len(args) > 2 {
		targetNodes = args[2:]
	} else {
		targetNodes = []string{"ZZZ"}
	}

	stepsRequired := 0
	var instructions string
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
	var stepsRequired *big.Int
	inputPath := args[0]

	var sourceNodes []string
	var targetNodes []string

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if line == "" || lineNum < 2 {
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

		switch source[len(source)-1] {
		case 'A':
			sourceNodes = append(sourceNodes, source)
		case 'Z':
			targetNodes = append(targetNodes, source)
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

	individualStepsRequired := make([]int, len(sourceNodes))
	for i, sourceNode := range sourceNodes {
		nArgs := append([]string{inputPath, sourceNode}, targetNodes...)
		stepsRequiredForOneString, err := Day8Part1(nArgs...)
		if err != nil {
			return "", err
		}

		stepsRequiredForOne, err := strconv.Atoi(stepsRequiredForOneString)
		if err != nil {
			return "", err
		}
		individualStepsRequired[i] = stepsRequiredForOne
	}

	gcd := func(a, b *big.Int) *big.Int {
		var temp big.Int
		return temp.GCD(nil, nil, a, b)
	}

	lcm := func(a, b *big.Int) *big.Int {
		var temp big.Int
		return temp.Div(temp.Mul(a, b), gcd(a, b))
	}

	stepsRequired = big.NewInt(int64(individualStepsRequired[0]))
	for i := 1; i < len(individualStepsRequired); i++ {
		stepsRequired = lcm(stepsRequired, big.NewInt(int64(individualStepsRequired[i])))
	}

	return stepsRequired.Text(10), nil
}
