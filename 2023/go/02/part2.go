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

func gamePower(gameSet [][]int) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	for _, set := range gameSet {
		if set[0] > maxRed {
			maxRed = set[0]
		}
		if set[1] > maxGreen {
			maxGreen = set[1]
		}
		if set[2] > maxBlue {
			maxBlue = set[2]
		}
	}

	return maxRed * maxGreen * maxBlue
}

func main() {
	lines := make([][][]int, 0, 100)
	absPath, _ := filepath.Abs("./02/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	redRe := regexp.MustCompile("[0-9]+ red")
	greenRe := regexp.MustCompile("[0-9]+ green")
	blueRe := regexp.MustCompile("[0-9]+ blue")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fullLine := strings.Split(scanner.Text(), ": ")[1]
		gameSet := make([][]int, 0, 20)
		for _, stringSet := range strings.Split(fullLine, "; ") {
			red := 0
			green := 0
			blue := 0
			redPart := redRe.FindString(stringSet)
			if redPart != "" {
				i, err := strconv.Atoi(strings.Split(redPart, " red")[0])
				if err != nil {
					panic(err)
				}
				red = i
			}
			greenPart := greenRe.FindString(stringSet)
			if greenPart != "" {
				i, err := strconv.Atoi(strings.Split(greenPart, " green")[0])
				if err != nil {
					panic(err)
				}
				green = i
			}
			bluePart := blueRe.FindString(stringSet)
			if bluePart != "" {
				i, err := strconv.Atoi(strings.Split(bluePart, " blue")[0])
				if err != nil {
					panic(err)
				}
				blue = i
			}
			gameSet = append(gameSet, []int{red, green, blue})
		}
		lines = append(lines, gameSet)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, line := range lines {
		total += gamePower(line)
	}

	fmt.Println(total)
}
