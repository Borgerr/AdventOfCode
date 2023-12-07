package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type handAndBid struct {
	hand string
	bid  int
}

func getHandsAndBids(lines []string) []handAndBid {
	handsAndBids := make([]handAndBid, 0, 1000)
	for _, line := range lines {
		spl := strings.Split(line, " ")

		hand := spl[0]
		bid, err := strconv.Atoi(spl[1])
		if err != nil {
			panic(err)
		}

		handsAndBids = append(handsAndBids, handAndBid{hand: hand, bid: bid})
	}
	return handsAndBids
}

func handToHex(hand string) int {
	// in hex:
	// 2 -> 2
	// 3 -> 3
	// ...
	// T -> a
	// J -> b
	// Q -> c
	// K -> d
	// A -> e
	returned := 0
	for _, char := range hand {
		if char == '2' {
			returned |= 0x2
		} else if char == '3' {
			returned |= 0x3
		} else if char == '4' {
			returned |= 0x4
		} else if char == '5' {
			returned |= 0x5
		} else if char == '6' {
			returned |= 0x6
		} else if char == '7' {
			returned |= 0x7
		} else if char == '8' {
			returned |= 0x8
		} else if char == '9' {
			returned |= 0x9
		} else if char == 'T' {
			returned |= 0xa
		} else if char == 'J' {
			returned |= 0xb
		} else if char == 'Q' {
			returned |= 0xc
		} else if char == 'K' {
			returned |= 0xd
		} else if char == 'A' {
			returned |= 0xe
		}

		returned <<= 4
	}
	return returned
}

func handDominates(h1 string, h2 string) bool {
	type1 := handType(h1)
	type2 := handType(h2)

	if type1 == type2 {
		return handToHex(h1) > handToHex(h2)
	} else {
		return type1 > type2
	}
}

func handType(hand string) int {
	cardCounts := []int{0, 0, 0, 0, 0}
	for index, comparedCard := range hand {
		for _, otherCard := range hand {
			if comparedCard == otherCard {
				cardCounts[index]++
			}
		}
	}

	var ty int
	if slices.Contains(cardCounts, 5) {
		ty = 7
	} else if slices.Contains(cardCounts, 4) {
		ty = 6
	} else if slices.Contains(cardCounts, 3) && slices.Contains(cardCounts, 2) {
		ty = 5
	} else if slices.Contains(cardCounts, 3) {
		ty = 4
	} else if slices.Contains(cardCounts, 2) {
		// check if this is a two pair or a one pair
		oneFound := false
		twoFound := false
		for _, count := range cardCounts {
			if count == 2 && !oneFound {
				oneFound = true
			} else if count == 2 && oneFound {
				twoFound = true
			}
		}
		if twoFound {
			ty = 3
		} else {
			ty = 2
		}
	} else {
		ty = 1
	}

	return ty
}

func getLines() []string {
	lines := make([]string, 0, 1000)
	absPath, _ := filepath.Abs("./07/sample.txt")
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
	return lines
}

func main() {
	lines := getLines()
	handsAndBids := getHandsAndBids(lines)

	// sort by ty
	sort.Slice(handsAndBids,
		func(i, j int) bool {
			return handDominates(handsAndBids[j].hand, handsAndBids[i].hand)
		})

	fmt.Println(handsAndBids)

	winnings := 0
	for rank, handAndBid := range handsAndBids {
		winnings += (rank + 1) * handAndBid.bid
	}
	fmt.Println(winnings)
}
