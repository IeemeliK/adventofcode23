package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    bytes, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    data := string(bytes)
    lines := strings.Split(data, "\n")
    for _, line := range lines {
        formattedString := stringToNumber(line)
        fmt.Printf("line: %v\n", formattedString)
    }
}

func parseNumbersFromString(string)  {
    
}

func stringToNumber(txt string) string {
    numberMap := map[string]string{
        "one": "1", 
        "two": "2", 
        "three": "3", 
        "four": "4", 
        "five": "5",
        "six": "6", 
        "seven": "7", 
        "eight": "8", 
        "nine": "9", 
    }

    for key, value := range numberMap {
       if strings.Contains(txt, key) {
           strings.ReplaceAll(txt, key, value)
       } 
    }
    return txt
}
