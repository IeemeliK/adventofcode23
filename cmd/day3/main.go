package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../inputs/day3_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	res := parseLines(lines)
	fmt.Println(res)
	var sum int
	for _, v := range res {
		sum += v
	}
	fmt.Println(sum)
}
