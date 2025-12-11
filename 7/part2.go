package main

func part2(lines []string) int {
	status := make([]int, len(lines[0]))
	for i := range lines {
		for j := range lines[i] {
			c := lines[i][j]
			if c == s {
				status[j] = 1
				break
			}
			if c == hat {
				if status[j] > 0 {
					if j > 0 {
						status[j-1] = status[j] + status[j-1]
					}
					if j < len(lines[0])-1 {
						status[j+1] = status[j] + status[j+1]
					}
				}
				status[j] = 0
			}
		}
	}

	sum := 0
	for _, n := range status {
		sum += n
	}
	return sum
}
