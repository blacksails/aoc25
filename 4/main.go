package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rows := parseInput(input)

	nRemoved := 0
	for {
		n := countAccessibleRolls(rows)
		if n == 0 {
			break
		}
		toBeRemoved := make([][2]int, n)
		for i := range rows {
			for j := range rows[i] {
				if rows[i][j] != roll {
					continue
				}
				if isRollAccessible(rows, i, j) {
					toBeRemoved = append(toBeRemoved, [2]int{i, j})
				}
			}
		}
		for _, pos := range toBeRemoved {
			rows[pos[0]][pos[1]] = '.'
		}
		nRemoved += n
	}

	for _, row := range rows {
		fmt.Println(string(row))
	}
	fmt.Println("Number of rolls removed:", nRemoved)
}

func countAccessibleRolls(rows [][]rune) int {
	rollsAccessible := 0
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] != roll {
				continue
			}
			if isRollAccessible(rows, i, j) {
				rollsAccessible++
			}
		}
	}
	return rollsAccessible
}

func parseInput(input string) [][]rune {
	rowStrs := strings.Split(strings.TrimSpace(input), "\n")
	rows := make([][]rune, len(rowStrs))
	for i, rowStr := range rowStrs {
		rows[i] = []rune(rowStr)
	}
	return rows
}

const roll = '@'

func isRollAccessible(rows [][]rune, i, j int) bool {
	adjacentRolls := 0
	if !(i == 0) {
		if !(j == 0) {
			if rows[i-1][j-1] == roll {
				adjacentRolls++
			}
		}
		if rows[i-1][j] == roll {
			adjacentRolls++
		}
		if !(j == len(rows[i])-1) {
			if rows[i-1][j+1] == roll {
				adjacentRolls++
			}
		}
	}
	if !(j == 0) {
		if rows[i][j-1] == roll {
			adjacentRolls++
		}
	}
	if !(j == len(rows[i])-1) {
		if rows[i][j+1] == roll {
			adjacentRolls++
		}
	}

	if !(i == len(rows)-1) {
		if !(j == 0) {
			if rows[i+1][j-1] == roll {
				adjacentRolls++
			}
		}
		if rows[i+1][j] == roll {
			adjacentRolls++
		}
		if !(j == len(rows[i])-1) {
			if rows[i+1][j+1] == roll {
				adjacentRolls++
			}
		}
	}

	return adjacentRolls < 4
}
