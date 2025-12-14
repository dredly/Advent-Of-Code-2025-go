package day1

import (
	"fmt"
	"strings"

	"github.com/dredly/aoc2025/internal/files"
	"github.com/dredly/aoc2025/internal/maths"
	"github.com/dredly/aoc2025/parsing"
)

func Part2Answer() {
	input := files.MustRead("inputdata/day1/real.txt")
	fmt.Printf("Day 1 Part 2 answer: %d\n", password2(input))
}

func password2(input string) int {
	position := 50
	zeroOccurences := 0
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		newPosition, clicks := rotate(l, position)
		position = newPosition
		zeroOccurences += clicks
	}
	return zeroOccurences
}

func rotate(line string, position int) (int, int) {
	direction := line[0:1]
	magnitude := parsing.MustParseInt(line[1:])
	if direction == "L" {
		return rotateLeft(magnitude, position)
	}
	return rotateRight(magnitude, position)
}

func rotateLeft(magnitude, position int) (int, int) {
	newPosition := position - magnitude
	clicks := maths.Abs(newPosition / 100) 
	newPosition = newPosition % 100
	if newPosition == 0 {
		clicks++
	}
	if newPosition < 0 {
		if position != 0 {
			clicks++
		}
		newPosition = 100 + newPosition
	}
	return newPosition, clicks
}

func rotateRight(magnitude, position int) (int, int) {
	newPosition := position + magnitude
	clicks := newPosition / 100
	newPosition = newPosition % 100
	return newPosition, clicks
}