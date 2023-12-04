package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func getWinnersAndHaves(line string) ([]int, []int) {
	line = strings.Split(line, ": ")[1]
	l := strings.Split(line, " | ")
	winnersString := l[0]
	havesString := l[1]

	winnersSubstrings := strings.Split(winnersString, " ")
	havesSubstrings := strings.Split(havesString, " ")

	winnersInts := make([]int, 0, len(winnersSubstrings))
	havesInts := make([]int, 0, len(havesSubstrings))

	for _, winner := range winnersSubstrings {
		if winner == "" {
			continue
		}
		i, err := strconv.Atoi(winner)
		if err != nil {
			panic(err)
		}
		winnersInts = append(winnersInts, i)
	}
	for _, have := range havesSubstrings {
		if have == "" {
			continue
		}
		i, err := strconv.Atoi(have)
		if err != nil {
			panic(err)
		}
		havesInts = append(havesInts, i)
	}

	return winnersInts, havesInts
}

func getMatchCount(line string) int {
	winners, haves := getWinnersAndHaves(line)
	matchCount := 0
	for _, have := range haves {
		if slices.Contains(winners, have) {
			matchCount++
		}
	}
	return matchCount
}

func main() {
	lines := make([]string, 0, 220)
	absPath, _ := filepath.Abs("./04/input.txt")
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

	cardCopies := make([]int, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		cardCopies = append(cardCopies, 1) // each card starts with one copy
	}

	for index, line := range lines {
		matchCount := getMatchCount(line)
		for i := index + 1; i <= index+matchCount; i++ {
			cardCopies[i] += cardCopies[index]
		}
	}

	total := 0
	for _, cardCount := range cardCopies {
		total += cardCount
	}

	//fmt.Println(cardCopies)
	fmt.Println(total)
}
