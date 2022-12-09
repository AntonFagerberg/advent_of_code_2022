package day09

import (
	"fmt"
	"strconv"
	"strings"
)

func solve(input []byte, size int) int {
	xs := make([]int, size)
	ys := make([]int, size)
	res := make(map[string]bool, 0)

	xd, yd := 0, 0

	for _, line := range strings.Split(string(input), "\n") {
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "U":
			xd, yd = 0, 1
		case "D":
			xd, yd = 0, -1
		case "L":
			xd, yd = -1, 0
		case "R":
			xd, yd = 1, 0
		}

		for i, _ := strconv.Atoi(parts[1]); i > 0; i-- {
			xs[0], ys[0] = xs[0]+xd, ys[0]+yd

			for i := 0; i < size-1; i++ {
				x, y := xs[i], ys[i]
				xx, yy := xs[i+1], ys[i+1]
				xp, yp := pos(x-xx), pos(y-yy)

				if xp == 2 && yp == 0 {
					xs[i+1] += (x - xx) / 2
				} else if xp == 0 && yp == 2 {
					ys[i+1] += (y - yy) / 2
				} else if xp+yp > 2 {
					if x > xx {
						xs[i+1] = xx + 1
					} else {
						xs[i+1] = xx - 1
					}

					if y > yy {
						ys[i+1] = yy + 1
					} else {
						ys[i+1] = yy - 1
					}
				}
			}

			res[fmt.Sprint(xs[size-1])+"|"+fmt.Sprint(ys[size-1])] = true
		}
	}

	return len(res)
}

func pos(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
