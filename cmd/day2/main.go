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
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		id, valid := test(scanner.Text())
		if valid {
			sum += id
		}
	}
	fmt.Println(sum)
}

func test(game string) (int, bool) {
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

	valid := false
	for _, set := range sets {
		var greenTotal, redTotal, blueTotal int
		for _, match := range blueRegex.FindAllStringSubmatch(set, -1) {
			if str, err := strconv.Atoi(match[1]); err == nil {
				blueTotal += str
			}
		}

		for _, match := range greenRegex.FindAllStringSubmatch(set, -1) {
			if str, err := strconv.Atoi(match[1]); err == nil {
				greenTotal += str
			}
		}

		for _, match := range redRegex.FindAllStringSubmatch(set, -1) {
			if str, err := strconv.Atoi(match[1]); err == nil {
				redTotal += str
			}
		}
		valid = (blueTotal <= cubeCounts["blue"] && greenTotal <= cubeCounts["green"] && redTotal <= cubeCounts["red"])
		if !valid {
			break
		}
	}

	return id, valid
}
