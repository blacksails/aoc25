package main

func part1(lines []string) int {
	result := 0
	status := make([]int, len(lines[0]))
	for i := range lines {
		for j := range lines[i] {
			c := lines[i][j]
			if c == s {
				status[j] = 1
				break
			}
			if c == hat {
				if status[j] == 1 {
					result++
					status[j] = 0
				}
				if j > 0 {
					status[j-1] = 1
				}
				if j < len(lines[0])-1 {
					status[j+1] = 1
				}
			}
		}
	}
	return result
}

