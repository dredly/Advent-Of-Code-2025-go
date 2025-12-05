package day1

import (
	"fmt"
	"strings"

	"github.com/dredly/aoc2025/internal/files"
	"github.com/dredly/aoc2025/parsing"
)

func Part1Answer() {
	input := files.MustRead("inputdata/day1/real.txt")
	fmt.Printf("Day 1 Part 1 answer: %d\n", password(input))
}

func password(input string) int {
	position := 50
	zeroOccurences := 0
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		position = (position + rotationRight(l)) % 100
		if position == 0 {
			zeroOccurences++
		}
	}
	return zeroOccurences
}

func rotationRight(line string) int {
	direction := line[0:1]
	magnitude := parsing.MustParseInt(line[1:])
	if direction == "L" {
		return 100 - magnitude
	}
	return magnitude
}