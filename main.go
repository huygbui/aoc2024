package main

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)



func main() {
// 	input := `MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX
// `
	input, _ := utils.GetInput(4)
	lines := strings.Fields(input)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	fmt.Println(solve2(grid))
}

func solve1(grid [][]rune) int {
	count := 0

	for i, row := range grid{
		for j := range row {
			if grid[i][j] == 'X'{

				//check horizontal forward
				if j < len(row) - 3{
					if grid[i][j+1] == 'M' {
						if grid[i][j+2] == 'A' {
							if grid[i][j+3] == 'S' {
								count += 1
							}
						}
					}
				}
				
				//check horizontal backwards
				if j > 2{
					if grid[i][j-1] == 'M' {
						if grid[i][j-2] == 'A' {
							if grid[i][j-3] == 'S' {
								count += 1
							}
						}
					}
				}
	
				//check vertical forward
				if i < len(grid) - 3{
					if grid[i+1][j] == 'M' {
						if grid[i+2][j] == 'A' {
							if grid[i+3][j] == 'S' {
								count += 1
							}
						}
					}
				}

				//check vertical backward
				if i > 2{
					if grid[i-1][j] == 'M' {
						if grid[i-2][j] == 'A' {
							if grid[i-3][j] == 'S' {
								count += 1
							}
						}
					}
				}

				//check diagonal forward 1
				if i < len(grid) - 3 && j < len(row) - 3{
					if grid[i+1][j+1] == 'M' {
						if grid[i+2][j+2] == 'A' {
							if grid[i+3][j+3] == 'S' {
								count += 1
							}
						}
					}
				}

				//check diagonal forward 2
				if i > 2 && j < len(row) - 3{
					if grid[i-1][j+1] == 'M' {
						if grid[i-2][j+2] == 'A' {
							if grid[i-3][j+3] == 'S' {
								count += 1
							}
						}
					}
				}

				//check diagonal backward up
				if i > 2 && j > 2{
					if grid[i-1][j-1] == 'M' {
						if grid[i-2][j-2] == 'A' {
							if grid[i-3][j-3] == 'S' {
								count += 1
							}
						}
					}
				}

				//check diagonal backward down
				if i < len(grid) - 3 && j > 2{
					if grid[i+1][j-1] == 'M' {
						if grid[i+2][j-2] == 'A' {
							if grid[i+3][j-3] == 'S' {
								count += 1
							}
						}
					}
				}
			}
		}
	}

	return count
}


func solve2(grid [][]rune) int {
	count := 0

	for i, row := range grid{
		for j := range row {
			if grid[i][j] == 'A' && (i > 0 && j > 0 && i < len(grid) - 1 && j < len(row) - 1)  {
				diagonalCount := 0 
				if (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M') || (grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S'){
					diagonalCount += 1
				}

				if (grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M') || (grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S'){
					diagonalCount += 1
				}

				if diagonalCount == 2 {
					count += 1
				}
			} 
		}
	}

	return count
}
