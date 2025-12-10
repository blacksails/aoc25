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
	ranges := strings.Split(input, ",")
	var sum int
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic(err)
		}

		for n := start; n <= end; n++ {
			s := strconv.Itoa(n)
			l := len(s)
			if l == 1 {
				continue
			}

			for lss := 1; lss <= l/2; lss++ {
				if l%lss != 0 {
					continue
				}

				nss := l / lss
				allEqual := true
				for i := 1; i < nss; i++ {
					if s[i*lss:(i+1)*lss] != s[0:lss] {
						allEqual = false
						break
					}
				}

				if allEqual {
					sum += n
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
