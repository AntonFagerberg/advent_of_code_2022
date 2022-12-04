package day04

import (
	"strconv"
	"strings"
)

func part1(input []byte) int {
	score := 0
	for _, l := range strings.Split(string(input), "\n") {
		parts := strings.Split(l, ",")
		first := strings.Split(parts[0], "-")
		second := strings.Split(parts[1], "-")
		a, _ := strconv.Atoi(first[0])
		b, _ := strconv.Atoi(first[1])
		x, _ := strconv.Atoi(second[0])
		y, _ := strconv.Atoi(second[1])

		if (a >= x && b <= y) || (x >= a && y <= b) {
			score += 1
		}
	}

	return score
}

func part2(input []byte) int {
	score := 0
	for _, l := range strings.Split(string(input), "\n") {
		parts := strings.Split(l, ",")
		first := strings.Split(parts[0], "-")
		second := strings.Split(parts[1], "-")
		a, _ := strconv.Atoi(first[0])
		b, _ := strconv.Atoi(first[1])
		x, _ := strconv.Atoi(second[0])
		y, _ := strconv.Atoi(second[1])

		if (a >= x && a <= y) ||
			(b >= x && b <= y) ||
			(x >= a && x <= b) ||
			(y >= a && y <= b) {
			score += 1
		}
	}

	return score
}
