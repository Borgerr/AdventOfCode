package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func getMinWithIndex(array []int) (int, int) {
	min := array[0]
	minIndex := 0
	for i := 1; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
			minIndex = i
		}
	}
	return min, minIndex
}

func getMaxWithIndex(array []int) (int, int) {
	max := array[0]
	maxIndex := 0
	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
			maxIndex = i
		}
	}
	return max, maxIndex
}

func processLine(line string) int {
	returned := 0
	// find earliest digit first
	digitIndexes := [9]int{
		MaxInt, MaxInt, MaxInt,
		MaxInt, MaxInt, MaxInt,
		MaxInt, MaxInt, MaxInt}
	digits := [9]string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	wordIndexes := digitIndexes
	words := [9]string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for index, word := range words {
		earliest := strings.Index(line, word)
		if earliest != -1 {
			wordIndexes[index] = earliest
		}
	}

	for index, digit := range digits {
		earliest := strings.Index(line, digit)
		if earliest != -1 {
			digitIndexes[index] = earliest
		}
	}

	earliestDigit, earliestDigIndex := getMinWithIndex(digitIndexes[:])
	earliestWord, earliestWordIndex := getMinWithIndex(wordIndexes[:])

	if earliestDigit < earliestWord {
		returned += (earliestDigIndex + 1) * 10
	} else {
		returned += (earliestWordIndex + 1) * 10
	}

	// do the same and find latest digit
	digitIndexes = [9]int{
		MinInt, MinInt, MinInt,
		MinInt, MinInt, MinInt,
		MinInt, MinInt, MinInt}
	wordIndexes = digitIndexes

	for index, word := range words {
		latest := strings.LastIndex(line, word)
		if latest != -1 {
			wordIndexes[index] = latest
		}
	}

	for index, digit := range digits {
		latest := strings.LastIndex(line, digit)
		if latest != -1 {
			wordIndexes[index] = latest
		}
	}

	latestDigit, latestDigIndex := getMaxWithIndex(digitIndexes[:])
	latestWord, latestWordIndex := getMaxWithIndex(wordIndexes[:])

	if latestDigit > latestWord {
		returned += latestDigIndex + 1
	} else {
		returned += latestWordIndex + 1
	}

	return returned
}

func main() {
	lines := make([]string, 0, 1000)
	absPath, _ := filepath.Abs("./01/input.txt")
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
