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

type gears struct {
	gearOne int
	gearTwo int
}

func findSymbolPositions(strings []string) map[position]gears {
	returned := make(map[position]gears)
	re := regexp.MustCompile(`\*`)
	for y, str := range strings {
		for _, l := range re.FindAllStringIndex(str, -1) {
			x := l[0]
			returned[position{x: x, y: y}] = gears{gearOne: 0, gearTwo: 0}
		}
	}
	return returned
}

func getPositions(positionsMap map[position]gears) []position {
	positions := make([]position, 0, len(positionsMap))
	for k := range positionsMap {
		positions = append(positions, k)
	}
	return positions
}

func getTotal(strings []string, positionsMap map[position]gears) int {
	re := regexp.MustCompile(`[0-9]+`)
	positions := getPositions(positionsMap)
	for y, str := range strings {
		for _, l := range re.FindAllStringIndex(str, -1) {
			i, err := strconv.Atoi(str[l[0]:l[1]])
			if err != nil {
				panic(err)
			}

			for x := l[0] - 1; x < l[1]+1; x++ {
				lines := [3]position{
					{x: x, y: y - 1},
					{x: x, y: y},
					{x: x, y: y + 1},
				}
				for _, line := range lines {
					if slices.Contains(positions, line) {
						gearInstance := positionsMap[line]
						if gearInstance.gearOne == 0 {
							positionsMap[line] = gears{gearOne: i, gearTwo: 0}
						} else if gearInstance.gearTwo == 0 {
							positionsMap[line] = gears{gearOne: gearInstance.gearOne, gearTwo: i}
						}
					}
				}
			}
		}
	}
	total := 0
	for k := range positionsMap {
		total += positionsMap[k].gearOne * positionsMap[k].gearTwo
	}
	return total
}

func main() {
	lines := make([]string, 0, 140)
	absPath, _ := filepath.Abs("./03/input.txt")
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
