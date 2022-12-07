package day07

import (
	"strconv"
	"strings"
)

func parse(input []byte) (map[string]int, map[string]int) {
	path := ""
	size := make(map[string]int, 0)
	paths := make(map[string]int, 0)

	for _, line := range strings.Split(string(input), "\n") {
		cmd := strings.Split(line, " ")
		if cmd[0] == "$" {
			if cmd[1] == "cd" {
				switch cmd[2] {
				case "/":
					path = "/"
				case "..":
					parts := strings.Split(path, "/")
					path = "/"
					for _, p := range parts[0 : len(parts)-2] {
						if p != "" {
							path += p + "/"
						}
					}
				default:
					path = path + cmd[2] + "/"
					paths[path] = 0
				}
			}
		} else if cmd[0] != "dir" {
			i, _ := strconv.Atoi(cmd[0])
			size[path+cmd[1]] = i
		}
	}

	for d := range paths {
		for f, i := range size {
			if strings.HasPrefix(f, d) {
				paths[d] += i
			}
		}
	}

	return paths, size
}

func part1(input []byte) int {
	paths, _ := parse(input)

	sum := 0

	for _, i := range paths {
		if 100000-i > 0 {
			sum += i
		}
	}

	return sum
}

func part2(input []byte) int {
	paths, size := parse(input)

	now := 0
	for _, i := range size {
		now += i
	}

	res := -1
	for _, i := range paths {
		if (i < res || res == -1) && i >= now+30000000-70000000 {
			res = i
		}
	}

	return res
}
