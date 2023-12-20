// Day 5: If You Give A Seed A Fertilizer
// https://adventofcode.com/2023/day/5

package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day5Part1(args ...string) (string, error) {
	// What is the lowest location number that corresponds to any of the initial seed numbers?
	inputPath := args[0]
	lowestLocationNumber := -1

	seedToSoil := "seed-to-soil"
	soilToFertilizer := "soil-to-fertilizer"
	fertilizerToWater := "fertilizer-to-water"
	waterToLight := "water-to-light"
	lightToTemperature := "light-to-temperature"
	temperatureToHumidity := "temperature-to-humidity"
	humidityToLocation := "humidity-to-location"

	almanacMapNames := []string{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight,
		lightToTemperature, temperatureToHumidity, humidityToLocation}
	almanacMaps := make(map[string][][3]int)

	addToAlmanacMap := func(mapName string, line string) error {
		var mapValues [3]int

		for i, mapValueString := range strings.Split(line, " ") {
			mapValue, err := strconv.Atoi(mapValueString)
			if err != nil {
				return err
			}

			mapValues[i] = mapValue
		}

		almanacMaps[mapName] = append(almanacMaps[mapName], mapValues)

		return nil
	}

	getDestinationNumber := func(sourceNumber int, mapName string) (destinationNumber int) {
		destinationNumber = sourceNumber

		for _, mapRange := range almanacMaps[mapName] {
			dStart, sStart, steps := mapRange[0], mapRange[1], mapRange[2]-1
			_, sEnd := dStart+steps, sStart+steps

			if sourceNumber >= sStart && sourceNumber <= sEnd {
				destinationNumber = dStart - sStart + sourceNumber
			}
		}

		return destinationNumber
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var seedValues []int
	onMap := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		goNextLine := false // :(

		if strings.Contains(line, "seeds:") {
			for _, seedValueString := range strings.Split(line[7:], " ") {
				seedValue, err := strconv.Atoi(seedValueString)
				if err != nil {
					return "", err
				}
				seedValues = append(seedValues, seedValue)
			}
			goNextLine = true
		}

		for _, almanacMapName := range almanacMapNames {
			if line == almanacMapName+" map:" {
				onMap = almanacMapName
				goNextLine = true
			}
		}

		if line == "" {
			goNextLine = true
			onMap = ""
		}
		if goNextLine {
			continue
		}

		err = addToAlmanacMap(onMap, line)
		if err != nil {
			return "", err
		}
	}

	for _, seedValue := range seedValues {
		soilValue := getDestinationNumber(seedValue, seedToSoil)
		fertilizerValue := getDestinationNumber(soilValue, soilToFertilizer)
		waterValue := getDestinationNumber(fertilizerValue, fertilizerToWater)
		lightValue := getDestinationNumber(waterValue, waterToLight)
		temperatureValue := getDestinationNumber(lightValue, lightToTemperature)
		humidityValue := getDestinationNumber(temperatureValue, temperatureToHumidity)
		locationValue := getDestinationNumber(humidityValue, humidityToLocation)

		if locationValue < lowestLocationNumber || lowestLocationNumber == -1 {
			lowestLocationNumber = locationValue
		}
	}

	return strconv.Itoa(lowestLocationNumber), nil
}

func Day5Part2(args ...string) (string, error) {
	// What is the lowest location number that corresponds to any of the initial seed numbers?
	inputPath := args[0]
	lowestLocationNumber := -1

	seedToSoil := "seed-to-soil"
	soilToFertilizer := "soil-to-fertilizer"
	fertilizerToWater := "fertilizer-to-water"
	waterToLight := "water-to-light"
	lightToTemperature := "light-to-temperature"
	temperatureToHumidity := "temperature-to-humidity"
	humidityToLocation := "humidity-to-location"

	almanacMapNames := []string{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight,
		lightToTemperature, temperatureToHumidity, humidityToLocation}
	almanacMaps := make(map[string][][3]int)

	addToAlmanacMap := func(mapName string, line string) error {
		var mapValues [3]int

		for i, mapValueString := range strings.Split(line, " ") {
			mapValue, err := strconv.Atoi(mapValueString)
			if err != nil {
				return err
			}

			mapValues[i] = mapValue
		}

		almanacMaps[mapName] = append(almanacMaps[mapName], mapValues)

		return nil
	}

	getDestinationNumber := func(sourceNumber int, mapName string) (destinationNumber int) {
		destinationNumber = sourceNumber

		for _, mapRange := range almanacMaps[mapName] {
			dStart, sStart, steps := mapRange[0], mapRange[1], mapRange[2]-1
			_, sEnd := dStart+steps, sStart+steps

			if sourceNumber >= sStart && sourceNumber <= sEnd {
				destinationNumber = dStart - sStart + sourceNumber
				break
			}
		}

		return destinationNumber
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var seedValueRanges []int
	onMap := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		goNextLine := false // :(

		if strings.Contains(line, "seeds:") {
			for _, seedValueString := range strings.Split(line[7:], " ") {
				seedValue, err := strconv.Atoi(seedValueString)
				if err != nil {
					return "", err
				}
				seedValueRanges = append(seedValueRanges, seedValue)
			}
			goNextLine = true
		}

		for _, almanacMapName := range almanacMapNames {
			if line == almanacMapName+" map:" {
				onMap = almanacMapName
				goNextLine = true
			}
		}

		if line == "" {
			goNextLine = true
			onMap = ""
		}
		if goNextLine {
			continue
		}

		err = addToAlmanacMap(onMap, line)
		if err != nil {
			return "", err
		}
	}

	for i := 0; i < len(seedValueRanges); i += 2 {
		start, size := seedValueRanges[i], seedValueRanges[i+1]
		for seedValue := start; seedValue < start+size; seedValue++ {
			soilValue := getDestinationNumber(seedValue, seedToSoil)
			fertilizerValue := getDestinationNumber(soilValue, soilToFertilizer)
			waterValue := getDestinationNumber(fertilizerValue, fertilizerToWater)
			lightValue := getDestinationNumber(waterValue, waterToLight)
			temperatureValue := getDestinationNumber(lightValue, lightToTemperature)
			humidityValue := getDestinationNumber(temperatureValue, temperatureToHumidity)
			locationValue := getDestinationNumber(humidityValue, humidityToLocation)

			if locationValue < lowestLocationNumber || lowestLocationNumber == -1 {
				lowestLocationNumber = locationValue
			}
		}
	}

	return strconv.Itoa(lowestLocationNumber), nil
}
