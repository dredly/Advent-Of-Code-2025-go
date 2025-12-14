package day2

import (
	"fmt"
	"slices"

	"github.com/dredly/aoc2025/internal/files"
)

func Part2Answer() {
	input := files.MustRead("inputdata/day2/real.txt")
	fmt.Printf("Day 2 Part 2 answer: %d\n", sumInvalidIDs(input, checkRangePart2))
}

func checkRangePart2(start, end int) map[int]struct{} {
	hits := make(map[int]struct{})
	for i := start; i <= end; i++ {
		d := intToDigits(i)
		for j := 1; j <= len(d) / 2; j++ {
			if len(d) % j != 0 {
				continue
			}
			equalSlcs := slices.Collect(slices.Chunk(d, j))
			if allSlicesEqual(equalSlcs) {
				hits[i] = struct{}{}
				break
			}
		}
	}
	return hits
}

func getFactors(n int) []int {
	return []int{}
}

func allSlicesEqual[T comparable](slcs [][]T) bool {
	if len(slcs) <= 1 {
		return true
	}
	for i := 1; i < len(slcs); i++ {
		if !slices.Equal(slcs[i -1], slcs[i]) {
			return false
		}
	}
	return true
}