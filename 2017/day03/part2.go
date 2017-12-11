package main

import (
	"fmt"
	"os"
	"strconv"
)

func expand_matrix(matrix [][]int) [][]int {
	size := len(matrix[0]) + 2

	var new_matrix [][]int
	for i := 0; i < size; i += 1 {
		if i == 0 || i == size-1 {
			new_matrix = append(new_matrix, make([]int, size))
			continue
		}
		new_row := []int{0}
		new_row = append(new_row, matrix[i-1]...)
		new_row = append(new_row, 0)
		new_matrix = append(new_matrix, new_row)
	}
	return new_matrix
}

var positions = [8][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func fill_position(matrix [][]int, row, col int) [][]int {
	sum := 0
	for _, pos := range positions {
		new_row := row + pos[0]
		new_col := col + pos[1]
		if new_row >= 0 && new_col >= 0 && new_row < len(matrix) && new_col < len(matrix) {
			sum += matrix[new_row][new_col]
		}
	}
	matrix[row][col] = sum
	return matrix
}

func fill_matrix(matrix [][]int, row_start, col_start int) [][]int {
	for ; row_start > 0; row_start -= 1 {
		matrix = fill_position(matrix, row_start, col_start)
	}
	for ; col_start > 0; col_start -= 1 {
		matrix = fill_position(matrix, row_start, col_start)
	}
	for ; row_start < len(matrix)-1; row_start += 1 {
		matrix = fill_position(matrix, row_start, col_start)
	}
	for ; col_start < len(matrix); col_start += 1 {
		matrix = fill_position(matrix, row_start, col_start)
	}
	return matrix
}

func find_answer(matrix [][]int, row_start, col_start, target int) int {
	for ; row_start > 0; row_start -= 1 {
		if matrix[row_start][col_start] > target {
			return matrix[row_start][col_start]
		}
	}
	for ; col_start > 0; col_start -= 1 {
		if matrix[row_start][col_start] > target {
			return matrix[row_start][col_start]
		}
	}
	for ; row_start < len(matrix)-1; row_start += 1 {
		if matrix[row_start][col_start] > target {
			return matrix[row_start][col_start]
		}
	}
	for ; col_start < len(matrix); col_start += 1 {
		if matrix[row_start][col_start] > target {
			return matrix[row_start][col_start]
		}
	}
	return 0
}

func main() {
	matrix := [][]int{{1}}

	target, _ := strconv.Atoi(os.Args[1])

	for {
		matrix = expand_matrix(matrix)
		matrix = fill_matrix(matrix, len(matrix)-2, len(matrix[0])-1)

		if matrix[len(matrix)-1][len(matrix)-1] > target {
			break
		}
	}

	// Print first number > target
	fmt.Println("Answer:", find_answer(matrix, len(matrix)-2, len(matrix[0])-1, target))

}
