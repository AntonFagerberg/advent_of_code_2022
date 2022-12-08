package day08

import (
	"strconv"
	"strings"
)

func part1(input []byte) int {
	strings := strings.Split(string(input), "\n")
	size := len(strings)
	grid := make([][]int, size)
	visible := make([][]bool, size)

	for x, line := range strings {
		grid[x] = make([]int, len(line))
		visible[x] = make([]bool, len(line))

		for y, v := range line {
			grid[x][y], _ = strconv.Atoi(string(v))
		}
	}

	for x := 0; x < size; x++ {
		prev := -1
		for y := 0; y < size; y++ {
			v := grid[x][y]
			if v > prev {
				visible[x][y] = true
				prev = v
			}
		}
	}

	for x := 0; x < size; x++ {
		prev := -1
		for y := size - 1; y >= 0; y-- {
			v := grid[x][y]
			if v > prev {
				visible[x][y] = true
				prev = v
			}
		}
	}

	for y := 0; y < size; y++ {
		prev := -1
		for x := 0; x < size; x++ {
			v := grid[x][y]
			if v > prev {
				visible[x][y] = true
				prev = v
			}
		}
	}

	for y := 0; y < size; y++ {
		prev := -1
		for x := size - 1; x >= 0; x-- {
			v := grid[x][y]
			if v > prev {
				visible[x][y] = true
				prev = v
			}
		}
	}

	sum := 0

	for _, l := range visible {
		for _, v := range l {
			if v {
				sum++
			}
		}
	}

	return sum
}

func part2(input []byte) int {
	strings := strings.Split(string(input), "\n")
	size := len(strings)
	grid := make([][]int, size)

	for x, line := range strings {
		grid[x] = make([]int, len(line))

		for y, v := range line {
			grid[x][y], _ = strconv.Atoi(string(v))
		}
	}

	max := -1

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			sum := 1
			h := grid[x][y]

			score := 0
			for xx := x + 1; xx < size; xx++ {
				score++
				if grid[xx][y] >= h {
					break
				}
			}
			sum *= score

			score = 0
			for xx := x - 1; xx >= 0; xx-- {
				score++
				if grid[xx][y] >= h {
					break
				}
			}
			sum *= score

			score = 0
			for yy := y + 1; yy < size; yy++ {
				score++
				if grid[x][yy] >= h {
					break
				}
			}
			sum *= score

			score = 0
			for yy := y - 1; yy >= 0; yy-- {
				score++
				if grid[x][yy] >= h {
					break
				}
			}
			sum *= score

			if sum > max {
				max = sum
			}
		}
	}

	return max
}
