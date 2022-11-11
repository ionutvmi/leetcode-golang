package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {

	chars := strings.Split(s, "")
	numColumns := len(s)

	matrix := make([][]string, numRows)
	for i := range matrix {
		matrix[i] = make([]string, numColumns)
	}

	row := 0
	column := 0
	direction := "down"

	for _, c := range chars {

		matrix[row][column] = c

		if direction == "down" {
			if row+1 == numRows {
				row = max(0, row-1)
				direction = "up"
				column += 1
			} else {
				row += 1
			}
		} else {
			if row-1 < 0 {
				direction = "down"

				if numRows > 1 {
					row = 1
				} else {
					column += 1
				}
			} else {
				row -= 1
				column += 1
			}
		}

	}

	result := ""

	for i := range matrix {
		for j := range matrix[i] {
			result += matrix[i][j]
		}
	}

	// printMatrix(matrix)

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func printMatrix(m [][]string) {

	for i := range m {
		for j := range m[i] {
			v := m[i][j]
			if v == "" {
				fmt.Print("-", " ")
			} else {
				fmt.Print(v, " ")
			}
		}
		fmt.Println("")
	}
}
