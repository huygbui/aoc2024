package day04

import "testing" 

func TestSolve(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	expected := 18
	result := Solve(input)

	if result != expected {
        t.Errorf("Expected %d XMAS occurrences, but got %d", expected, result)
    }
}