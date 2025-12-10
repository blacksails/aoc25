package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	ps, err := v1Problems(input)
	if err != nil {
		panic(err)
	}

	r, err := calculateProblems(ps)
	if err != nil {
		panic(err)
	}

	fmt.Printf("v1: %d\n", r)

	ps, err = v2Problems(input)
	if err != nil {
		panic(err)
	}

	r, err = calculateProblems(ps)
	if err != nil {
		panic(err)
	}

	fmt.Printf("v2: %d\n", r)
}

func calculateProblems(ps []mathProblem) (int, error) {
	s := 0
	for _, p := range ps {
		var r int
		switch p.operator {
		case add:
			for _, n := range p.numbers {
				r += n
			}
		case multiply:
			r = 1
			for _, n := range p.numbers {
				r *= n
			}
		default:
			return 0, fmt.Errorf("unknown operator: %c", p.operator)
		}
		fmt.Printf("result: %d\n", r)
		s += r
	}
	return s, nil
}

type mathProblem struct {
	numbers  []int
	operator operator
}

type operator byte

const (
	add      operator = '+'
	multiply operator = '*'
)
