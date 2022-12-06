package day06

func solve(input string, size int) int {
	for i := 0; i < len(input)-size-1; i++ {
		m := make(map[byte]bool)

		for j := i; j < i+size; j++ {
			m[input[j]] = true
		}

		if len(m) == size {
			return i + size
		}
	}

	return -1
}
