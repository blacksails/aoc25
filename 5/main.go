package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	numberOfFreshIngredients := 0
	ranges, ids := parseInput(input)
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				numberOfFreshIngredients++
				break
			}
		}
	}
	fmt.Println(numberOfFreshIngredients)

	totalFreshIngredients := 0
	normalizedRanges := normalizeRanges(ranges)
	for _, r := range normalizedRanges {
		totalFreshIngredients += r.end - r.start + 1
	}

	fmt.Println(totalFreshIngredients)
}

func normalizeRanges(ranges []ingredientRange) []ingredientRange {
	slices.SortFunc(ranges, func(a, b ingredientRange) int {
		if a.start != b.start {
			return a.start - b.start
		}
		return a.end - b.end
	})

	var result []ingredientRange
	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		if current.end >= r.start {
			if current.end < r.end {
				current.end = r.end
			}
		} else {
			result = append(result, current)
			current = r
		}
	}
	result = append(result, current)

	return result
}

type ingredientRange struct {
	start int
	end   int
}

func parseInput(input string) ([]ingredientRange, []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var ranges []ingredientRange
	var ids []int
	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, ingredientRange{start: start, end: end})
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		id, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	return ranges, ids
}
