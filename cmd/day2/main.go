package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum, powerSum int
	for scanner.Scan() {
		id, power, valid := test(scanner.Text())
		if valid {
			sum += id
		}
		powerSum += power
	}
	fmt.Printf("Sum: %d\nPower Sum: %d\n", sum, powerSum)
}

func test(game string) (int, int, bool) {
	blueRegex := regexp.MustCompile(`(\d+)(?:\sblue)`)
	greenRegex := regexp.MustCompile(`(\d+)(?:\sgreen)`)
	redRegex := regexp.MustCompile(`(\d+)(?:\sred)`)

	gameRegex := regexp.MustCompile(`Game\s(\d+)`)
	id, _ := strconv.Atoi(gameRegex.FindStringSubmatch(game)[1])

	cubeCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sets := strings.Split(game, ";")

	valid := true
	var redMax, blueMax, greenMax int
	for _, set := range sets {
		var greenTotal, redTotal, blueTotal int
		for _, match := range blueRegex.FindAllStringSubmatch(set, -1) {
			if count, err := strconv.Atoi(match[1]); err == nil {
				blueTotal += count
				blueMax = max(blueMax, count)
			}
		}

		for _, match := range greenRegex.FindAllStringSubmatch(set, -1) {
			if count, err := strconv.Atoi(match[1]); err == nil {
				greenTotal += count
				greenMax = max(greenMax, count)
			}
		}

		for _, match := range redRegex.FindAllStringSubmatch(set, -1) {
			if count, err := strconv.Atoi(match[1]); err == nil {
				redTotal += count
				redMax = max(redMax, count)
			}
		}
		valid = valid && greenTotal <= cubeCounts["green"] && redTotal <= cubeCounts["red"] && blueTotal <= cubeCounts["blue"]
	}

	power := redMax * blueMax * greenMax
	return id, power, valid
}
