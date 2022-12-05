package day05

import (
	"strconv"
	"strings"
)

func solve(input []byte, part int) string {
	stacks := make([][]rune, 0)
	cmds := make([][]int, 0)
	parseFirst := true

	for ii, line := range strings.Split(string(input), "\n") {
		if ii == 0 {
			for z := 0; z <= len(line)/4; z++ {
				stacks = append(stacks, make([]rune, 0))
			}
		}

		if line == "" {
			parseFirst = false

			for i, s := range stacks {
				stacks[i] = s[0 : len(s)-1]
			}
		} else if parseFirst {
			for i, c := range line {
				if c != '[' && c != ']' && c != ' ' {
					stacks[i/4] = append(stacks[i/4], c)
				}
			}
		} else {
			parts := strings.Split(line, " ")
			move, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])
			cmds = append(cmds, []int{move, from - 1, to - 1})
		}
	}

	for _, cmd := range cmds {
		v := stacks[cmd[1]][0:cmd[0]]

		switch part {
		case 1:
			for i := 0; i < len(v); i++ {
				stacks[cmd[2]] = append([]rune{v[i]}, stacks[cmd[2]]...)
			}
		case 2:
			for i := len(v) - 1; i >= 0; i-- {
				stacks[cmd[2]] = append([]rune{v[i]}, stacks[cmd[2]]...)
			}
		}

		stacks[cmd[1]] = stacks[cmd[1]][cmd[0]:len(stacks[cmd[1]])]
	}

	res := make([]rune, 0)
	for _, r := range stacks {
		res = append(res, r[0])
	}

	return string(res)
}
