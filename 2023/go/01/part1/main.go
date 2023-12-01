package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

func processLine(line string) int {
	digits := make([]rune, 0, 2)
	for _, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		}
	}

	i, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	lines := make([]string, 0, 1000)
	absPath, _ := filepath.Abs("./01/part1/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, line := range lines {
		total += processLine(line)
	}

	fmt.Printf("%d\n", total)
}
