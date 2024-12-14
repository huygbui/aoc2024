package main

import (
	"aoc2024/utils"
	"fmt"
)

func main() {
	inp, err := utils.GetInput(4)
	if err != nil {
        fmt.Printf("Error getting input: %v\n", err)
		return 
	}

    fmt.Printf("First 50 characters of input: %q\n", inp[:min(50,len(inp))])
}