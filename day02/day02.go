package day02

import "strings"

func parse(input []byte) [][]string {
	res := make([][]string, 0)

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		part := make([]string, 2)

		for i, item := range strings.Split(line, " ") {
			part[i] = item
		}

		res = append(res, part)
	}

	return res
}

// X,A = Rock
// Y,B = Paper
// Z,C = Sissor
func part1(input [][]string) int {
	score := 0

	for _, v := range input {
		switch v[1] {
		case "X":
			score += 1

			switch v[0] {
			case "A":
				score += 3
			case "B":
				score += 0
			case "C":
				score += 6
			}

		case "Y":
			score += 2

			switch v[0] {
			case "A":
				score += 6
			case "B":
				score += 3
			case "C":
				score += 0
			}

		case "Z":
			score += 3

			switch v[0] {
			case "A":
				score += 0
			case "B":
				score += 6
			case "C":
				score += 3
			}

		}
	}

	return score
}

// X,A = Rock,Lose
// Y,B = Paper,Draw
// Z,C = Sissor,Win
func part2(input [][]string) int {
	for _, v := range input {
		switch v[1] {
		case "X":
			switch v[0] {
			case "A":
				v[1] = "Z"
			case "B":
				v[1] = "X"
			case "C":
				v[1] = "Y"
			}

		case "Y":
			switch v[0] {
			case "A":
				v[1] = "X"
			case "B":
				v[1] = "Y"
			case "C":
				v[1] = "Z"
			}

		case "Z":
			switch v[0] {
			case "A":
				v[1] = "Y"
			case "B":
				v[1] = "Z"
			case "C":
				v[1] = "X"
			}

		}
	}

	return part1(input)
}
