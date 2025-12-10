package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	pos := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		steps, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil {
			panic(err)
		}

		switch direction {
		case 'R':
			password += get0Visits(pos, steps)
			pos = pos + steps
		case 'L':
			password += get0Visits(pos, -steps)
			pos = pos - steps
		default:
			panic("unknown direction: " + string(direction))
		}

		pos = pos % 100
		if pos < 0 {
			pos += 100
		}

		fmt.Printf("Direction: %c, Steps: %d, New Pos: %d, Password: %d\n", direction, steps, pos, password)
	}

	fmt.Println(password)
}

func get0Visits(pos int, steps int) int {
	if steps == 0 {
		return 0
	}

	if steps > 0 {
		return (pos + steps) / 100
	} else {
		stepsTo0 := pos
		if stepsTo0 == 0 {
			stepsTo0 = 100
		}

		remainingSteps := -steps - stepsTo0
		if remainingSteps < 0 {
			return 0
		}

		return 1 + remainingSteps/100
	}
}
