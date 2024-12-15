package main

import (
	"aoc2024/days/day04"
	"fmt"
)

func main() {
	res1, err := day04.Solve(1)
	if err != nil {
		fmt.Printf("Failed to solve part 1. Got: %v\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", *res1)
	}

	res2, err := day04.Solve(2)
	if err != nil {
		fmt.Printf("Failed to solve part 2. Got: %v\n", err)
	} else {
		fmt.Printf("Answer to Part 2 is %d\n", *res2)
	}
}

