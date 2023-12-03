package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
)

type position struct {
	x int
	y int
}

func findSymbolPositions(strings []string) []position {
	returned := make([]position, 0, 100000)
	re := regexp.MustCompile(`[\$@#%&\*-\+=/]`)
	for y, str := range strings {
		for _, l := range re.FindAllStringIndex(str, -1) {
			x := l[0]
			returned = append(returned, position{x: x, y: y})
		}
	}
	return returned
}

func getTotal(strings []string, positions []position) int {
	total := 0
	re := regexp.MustCompile(`[0-9]+`)
	for y, str := range strings {
		for _, l := range re.FindAllStringIndex(str, -1) {
			for x := l[0] - 1; x < l[1]+1; x++ {
				onPrevLine := slices.Contains(positions, position{x: x, y: y - 1})
				onThisLine := slices.Contains(positions, position{x: x, y: y})
				onNextLine := slices.Contains(positions, position{x: x, y: y + 1})
				if onPrevLine || onThisLine || onNextLine {
					i, err := strconv.Atoi(str[l[0]:l[1]])
					if err != nil {
						panic(err)
					}
					total += i
					break
				}
			}
		}
	}
	return total
}

func main() {
	lines := make([]string, 0, 140)
	absPath, _ := filepath.Abs("./03/sample.txt")
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
	positions := findSymbolPositions(lines)

	fmt.Println(getTotal(lines, positions))
}
