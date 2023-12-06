package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type compactRange struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLength      int
}

func getSeedNumbers(line string) []int {
	// seeds are on first line
	seedsString := strings.Split(line, "seeds: ")[1]
	seedSubstrings := strings.Split(seedsString, " ")
	seedNumbers := make([]int, 0, len(seedSubstrings))
	for _, seedNumberString := range seedSubstrings {
		i, err := strconv.Atoi(seedNumberString)
		if err != nil {
			panic(err)
		}
		seedNumbers = append(seedNumbers, i)
	}
	return seedNumbers
}

func getMap(lines []string) ([]compactRange, int) {
	returned := make([]compactRange, 100)

	mapLines := make([]string, 0, 20)
	var nextStart int
	for lineNum, line := range lines {
		if line == "" {
			nextStart = lineNum + 2
			break
		}
		mapLines = append(mapLines, line)
	}

	for _, mapLine := range mapLines {
		mapSubstrings := strings.Split(mapLine, " ")
		l := make([]int, 0, 3)
		for _, str := range mapSubstrings {
			i, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			l = append(l, i)
		}
		returned = append(returned, compactRange{destRangeStart: l[0], sourceRangeStart: l[1], rangeLength: l[2]})
	}

	return returned, nextStart
}

func lookup(val int, m *[]compactRange) int {
	for _, r := range *m {
		if (val >= r.sourceRangeStart) && (val < r.sourceRangeStart+r.rangeLength) {
			return r.destRangeStart + (val - r.sourceRangeStart)
		}
	}

	return val
}

func getLocations(seeds []int,
	sed2Soil *[]compactRange, soil2Fert *[]compactRange, fert2Wat *[]compactRange,
	wat2Lit *[]compactRange, lit2Temp *[]compactRange, temp2Hum *[]compactRange,
	hum2Loc *[]compactRange) []int {
	locations := make([]int, 0, len(seeds))

	var wg sync.WaitGroup

	for _, seed := range seeds {
		wg.Add(1)

		go func(s int) {
			defer wg.Done()

			soil := lookup(s, sed2Soil)
			fert := lookup(soil, soil2Fert)
			water := lookup(fert, fert2Wat)
			light := lookup(water, wat2Lit)
			temp := lookup(light, lit2Temp)
			hum := lookup(temp, temp2Hum)
			location := lookup(hum, hum2Loc)
			locations = append(locations, location)
		}(seed)
	}

	wg.Wait()
	return locations
}

func main() {
	lines := make([]string, 0, 255)
	absPath, _ := filepath.Abs("./05/input.txt")
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

	seedNums := getSeedNumbers(lines[0])

	lines = lines[3:]
	// line at index 0 is now line after "seed-to-soil map"
	seedToSoilMap, nextStart := getMap(lines)
	lines = lines[nextStart:]
	soilToFertMap, nextStart := getMap(lines)
	lines = lines[nextStart:]
	fertToWaterMap, nextStart := getMap(lines)
	lines = lines[nextStart:]
	waterToLightMap, nextStart := getMap(lines)
	lines = lines[nextStart:]
	lightToTempMap, nextStart := getMap(lines)
	lines = lines[nextStart:]
	tempToHumMap, nextStart := getMap(lines)
	lines = lines[nextStart:]
	humToLocMap, nextStart := getMap(lines)

	locations := getLocations(seedNums,
		&seedToSoilMap, &soilToFertMap, &fertToWaterMap,
		&waterToLightMap, &lightToTempMap, &tempToHumMap, &humToLocMap)

	minLoc := 0xfffffffffffffff
	for _, location := range locations {
		if location < minLoc {
			minLoc = location
		}
	}

	fmt.Println(minLoc)
}
