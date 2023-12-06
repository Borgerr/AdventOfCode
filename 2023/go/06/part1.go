package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func getNumsFromString(str string) []int {
	re := regexp.MustCompile(`[0-9]+`)
	intSubstrings := re.FindAllString(str, -1)

	returned := make([]int, 0, len(intSubstrings))
	for _, sub := range intSubstrings {
		i, err := strconv.Atoi(sub)
		if err != nil {
			panic(err)
		}
		returned = append(returned, i)
	}
	return returned
}

func getNumberOfWays(time int, currentRecord int) int {
	ways := 0
	for potentialHoldTime := 0; potentialHoldTime < time; potentialHoldTime++ {
		remainingTime := time - potentialHoldTime
		if remainingTime*potentialHoldTime > currentRecord {
			ways++
		}
	}
	return ways
}

func main() {
	lines := make([]string, 0, 2)
	absPath, _ := filepath.Abs("./06/input.txt")
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

	raceTimes := getNumsFromString(strings.Split(lines[0], "Time: ")[1])
	recordDists := getNumsFromString(strings.Split(lines[1], "Distance: ")[1])
	result := 1
	for raceNum := 0; raceNum < len(raceTimes); raceNum++ {
		result *= getNumberOfWays(raceTimes[raceNum], recordDists[raceNum])
	}
	fmt.Println(result)
}
