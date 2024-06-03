package main

func createMatrix(rows, cols int) [][]int {
	if rows == 0 || cols == 0 {
		return [][]int{}
	}

	var matrix [][]int

	for i := 0; i < rows; i++ {
		var currRow []int
		for j := 0; j < cols; j++ {
			currRow = append(currRow, i*j)
		}

		matrix = append(matrix, currRow)
	}

	return matrix
}
