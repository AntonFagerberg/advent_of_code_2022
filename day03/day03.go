package day03

import (
	"strings"
)

func score(r int) int {
	if r-'a' < 0 {
		return r - 'A' + 27
	} else {
		return r - 'a' + 1
	}
}

func part1(input []byte) int {
	lines := strings.Split(string(input), "\n")
	scores := make([]int, len(lines))

	for i, v := range lines {
		l := len(v)
		first := v[0 : l/2]
		last := v[l/2 : l]

		for _, a := range first {
			for _, b := range last {
				if a == b {
					scores[i] = score(int(a))
				}
			}
		}
	}

	sum := 0
	for _, v := range scores {
		sum += v
	}

	return sum
}

func part2(input []byte) int {
	lines := strings.Split(string(input), "\n")
	scores := make([]int, len(lines)/3)

	for i := 0; i < len(lines); i += 3 {
		for _, c1 := range lines[i] {
			for _, c2 := range lines[i+1] {
				for _, c3 := range lines[i+2] {
					if c1 == c2 && c2 == c3 {
						scores[i/3] = score(int(c1))
					}
				}
			}
		}
	}

	sum := 0
	for _, v := range scores {
		sum += v
	}

	return sum
}
