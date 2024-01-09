package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []string{line})
	}

	fmt.Println(matrix)
}

func test(matrix [][]string) {

}
