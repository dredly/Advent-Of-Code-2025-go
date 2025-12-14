package day2

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/dredly/aoc2025/internal/files"
	"github.com/dredly/aoc2025/parsing"
)

func Part1Answer() {
	input := files.MustRead("inputdata/day2/real.txt")
	fmt.Printf("Day 2 Part 1 answer: %d\n", sumInvalidIDs(input))
}

func sumInvalidIDs(input string) int {
	rangesRaw := strings.Split(input, ",")
	m := make(map[int]struct{})
	for _, r := range rangesRaw {
		nums := strings.Split(r, "-")
		start := parsing.MustParseInt(nums[0])
		end := parsing.MustParseInt(nums[1])
		hits := checkRange(start, end)
		m = mergeMaps(m, hits)
	}
	sum := 0
	for k := range m {
		sum += k
	}
	return sum
}

func checkRange(start, end int) map[int]struct{} {
	hits := make(map[int]struct{})
	for i := start; i <= end; i++ {
		d := intToDigits(i)
		if len(d) % 2 != 0 {
			continue
		}
		halfLen := len(d) / 2
		if slices.Equal(d[:halfLen], d[halfLen:]) {
			hits[i] = struct{}{}
		}

	}
	return hits
}

func mergeMaps[T comparable, U any](m1, m2 map[T]U) map[T]U {
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}

func intToDigits(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func digitsToInt(d []int) int {
	res := 0
	for i, digit := range d {
		multiplier := int(math.Pow10(len(d) - (i + 1)))
		res += multiplier * digit
	}
	return res
}

