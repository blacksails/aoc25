package main

import (
	"fmt"
	"strconv"
	"strings"
)

func v2Problems(input string) ([]mathProblem, error) {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	columns := makeLinesIntoColumns(lines)

	var problems []mathProblem
	var problem mathProblem
	var newProblem bool = false
	for i := range columns {
		// Reset problem if starting a new one
		if newProblem {
			problem = mathProblem{}
			newProblem = false
		}

		// Check if column is empty ensure next column starts a new problem
		columnIsEmpty := true
		for _, b := range columns[i] {
			if !isWhitespace(b) {
				columnIsEmpty = false
				break
			}
		}

		var numberStr string
		for _, b := range columns[i] {
			if isNumber(b) {
				numberStr += string(b)
				continue
			}
			if isOperator(b) {
				switch operator(b) {
				case add, multiply:
					problem.operator = operator(b)
				default:
					return nil, fmt.Errorf("unknown operator: %c", b)
				}
			}
		}

		if numberStr != "" {
			n, err := strconv.Atoi(numberStr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse number %s: %w", numberStr, err)
			}
			problem.numbers = append(problem.numbers, n)
		}

		if columnIsEmpty {
			problems = append(problems, problem)
			newProblem = true
			continue
		}
	}
	problems = append(problems, problem)
	return problems, nil
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isOperator(b byte) bool {
	return b == '+' || b == '*'
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n'
}

func makeLinesIntoColumns(lines []string) [][]byte {
	numColumns := len(lines[0])
	columns := make([][]byte, numColumns)
	for i := range numColumns {
		columns[i] = make([]byte, len(lines))
		for j := range lines {
			columns[i][j] = lines[j][i]
		}
	}
	return columns
}
