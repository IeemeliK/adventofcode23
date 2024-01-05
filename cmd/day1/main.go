package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    bytes, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    data := string(bytes)
    lines := strings.Split(data, "\n")
    total := 0
    for _, line := range lines {
        numberString := parseNumbersFromString(line)
        if number, err := strconv.Atoi(numberString); err == nil {
            fmt.Printf("Line: %s Number: %d\n", line, number)
            total += number
        }
    }
    fmt.Println(total)
}

func parseNumbersFromString(txt string) string {
    var first, last string
    for i := 0; i < len(txt); i++ {
        if _, err := strconv.Atoi(string(txt[i])); err == nil {
            first = string(txt[i])
            break
        }
    }

    for i := len(txt) - 1; i >= 0; i-- {
        if _, err := strconv.Atoi(string(txt[i])); err == nil {
            last = string(txt[i])
            break
        }
    }
    return first + last
}

// func stringToNumber(txt string) string {
//     numberMap := map[string]string{
//         "one": "1", 
//         "two": "2", 
//         "three": "3", 
//         "four": "4", 
//         "five": "5",
//         "six": "6", 
//         "seven": "7", 
//         "eight": "8", 
//         "nine": "9", 
//     }
//
//     for key, value := range numberMap {
//        if strings.Contains(txt, key) {
//            strings.ReplaceAll(txt, key, value)
//        } 
//     }
//     return txt
// }
