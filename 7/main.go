package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	s   = 'S'
	hat = '^'
)

func main() {
	lines := parseAsLines(input)

	fmt.Println("Activated splitters:", part1(lines))
	fmt.Println("Number of paths:", part2(lines))
}

func parseAsLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
