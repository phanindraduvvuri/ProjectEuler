package main

import "fmt"

// Maxtrix if a 2-D slice
type Maxtrix [][]int

const (
	gridSize int = 20
)

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}

	return val2

}

func manyMax(vals ...int) int {
	var maxVal int = -1

	for _, val := range vals {
		maxVal = max(val, maxVal)
	}

	return maxVal
}

func diaProd(grid Maxtrix) int {
	var maxProd int = -1

	for i := 0; i < gridSize-3; i++ {
		for j := 0; j < gridSize-3; j++ {
			prod := grid[i][j] * grid[i+1][j+1] * grid[i+2][i+2] * grid[i+3][i+3]

			maxProd = max(prod, maxProd)
		}
	}

	return maxProd
}

func revDiaProd(grid Maxtrix) int {
	var maxProd int = -1

	for i := 3; i < gridSize-3; i++ {
		for j := 3; j < gridSize-3; j++ {
			prod := grid[i][j] * grid[i+1][j-1] * grid[i+2][j-2] * grid[i+3][j-3]

			maxProd = max(prod, maxProd)
		}
	}

	return maxProd
}

func rightProd(grid Maxtrix) int {
	var maxProd int = -1

	for i := 0; i < gridSize-3; i++ {
		prod := grid[i][i] * grid[i][i+1] * grid[i][i+2] * grid[i][i+3]

		maxProd = max(prod, maxProd)
	}

	return maxProd
}

func downProd(grid Maxtrix) int {
	var maxProd int = -1

	for i := 0; i < gridSize-3; i++ {
		prod := grid[i][i] * grid[i+1][i] * grid[i+2][i] * grid[i+3][i]

		maxProd = max(prod, maxProd)
	}

	return maxProd
}

func main() {
	var temp, res int
	grid := make(Maxtrix, 20)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, 20)
	}

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			fmt.Scan(&temp)

			grid[i][j] = temp
		}
	}

	res = manyMax(diaProd(grid), revDiaProd(grid), rightProd(grid), downProd(grid))
	fmt.Println(res)
}
