package main

import (
	"fmt"
	"strconv"
	"strings"
)

func v1Problems(input string) ([]mathProblem, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	numberLines := make([][]int, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		fields := strings.Fields(lines[i])
		numberLines[i] = make([]int, len(fields))
		for j, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed parsing number %q: %w", field, err)
			}
			numberLines[i][j] = n
		}
	}

	var problems []mathProblem
	for field := range strings.FieldsSeq(lines[len(lines)-1]) {
		problems = append(problems, mathProblem{
			operator: operator(field[0]),
		})
	}

	for i := range numberLines {
		for j := range numberLines[i] {
			problems[j].numbers = append(problems[j].numbers, numberLines[i][j])
		}
	}

	return problems, nil
}
