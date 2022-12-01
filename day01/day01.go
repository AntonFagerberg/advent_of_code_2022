package day01

import (
	"sort"
	"strconv"
	"strings"
)

func Parse(input []byte) [][]int {
	res := make([][]int, 0)
	part := make([]int, 0)

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if line == "" {
			res = append(res, part)
			part = make([]int, 0)
		} else {
			i, _ := strconv.Atoi(line)
			// todo error check
			part = append(part, i)
		}
	}

	res = append(res, part)

	return res
}

func Sums(input [][]int) []int {
	res := make([]int, 0)

	for _, l := range input {
		sum := 0
		for _, v := range l {
			sum += v
		}

		res = append(res, sum)
	}

	sort.Ints(res)

	return res
}

func Part1(input []byte) int {
	res := Sums(Parse(input))
	return res[len(res)-1]
}

func Part2(input []byte) int {
	res := Sums(Parse(input))
	return res[len(res)-1] + res[len(res)-2] + res[len(res)-3]
}
