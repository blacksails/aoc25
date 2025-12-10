package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	banks, err := parseInput(input)
	if err != nil {
		panic(err)
	}

	totalJolt := 0
	for _, bank := range banks {
		if len(bank) < 12 {
			panic("bank must have at least 12 batteries")
		}

		const numJolts = 12
		maxJolt := make([]int, numJolts)
		for i := 0; i < 12; i++ {
			maxJolt[i], bank = maxWithAtLeastNDigitsAfter(bank, numJolts-i-1)
		}

		var maxJoltStr string
		for _, n := range maxJolt {
			maxJoltStr += strconv.Itoa(n)
		}

		maxJoltN, err := strconv.Atoi(maxJoltStr)
		if err != nil {
			panic(err)
		}
		totalJolt += maxJoltN
	}

	fmt.Println(totalJolt)
}

func maxWithAtLeastNDigitsAfter(digits []int, n int) (int, []int) {
	if len(digits) < n {
		return 0, nil
	}

	max := 0
	maxIndex := -1
	for i := 0; i < len(digits)-n; i++ {
		if max < digits[i] {
			max = digits[i]
			maxIndex = i
		}
	}
	return max, digits[maxIndex+1:]
}

func parseInput(input string) ([][]int, error) {
	bankStrs := strings.Split(input, "\n")

	var banks [][]int
	for _, bankStr := range bankStrs {
		bankStr := strings.TrimSpace(bankStr)
		if bankStr == "" {
			continue
		}

		bank := make([]int, len(bankStr))
		for j, n := range bankStr {
			joltRating, err := strconv.Atoi(string(n))
			if err != nil {
				return nil, err
			}
			bank[j] = joltRating
		}

		banks = append(banks, bank)
	}

	return banks, nil
}
