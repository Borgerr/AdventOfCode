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

func getNumFromString(str string) int {
	re := regexp.MustCompile(`[0-9]+`)
	intString := strings.Join(re.FindAllString(str, -1), "")

	i, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	return i
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

	raceTime := getNumFromString(strings.Split(lines[0], "Time: ")[1])
	recordDist := getNumFromString(strings.Split(lines[1], "Distance: ")[1])
	fmt.Println(getNumberOfWays(raceTime, recordDist))
}
