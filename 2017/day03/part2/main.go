package main

import "fmt"

const N int = 599

func fillCell(i, j int, grid [][]int) {
	di := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	dj := []int{0, 1, 1, 1, 0, -1, -1, -1}

	for x := range 8 {
		ix := i + di[x]
		jx := j + dj[x]
		grid[i][j] += grid[ix][jx]
	}

}

func getFirstBiggerVal(val int, grid [][]int) int {
	i := N / 2
	j := N / 2

	grid[i][j] = 1
	currSize := 0

	for {
		currSize += 2
		j++
		fillCell(i, j, grid)
		// fillCell(i, j, grid)
		// Move Up currSize - 1 times
		for range currSize - 1 {
			i--
			fillCell(i, j, grid)
			if grid[i][j] > val {
				return grid[i][j]
			}
		}
		// Move left currSize times
		for range currSize {
			j--
			fillCell(i, j, grid)
			if grid[i][j] > val {
				return grid[i][j]
			}
		}
		// Move down currSize times
		for range currSize {
			i++
			fillCell(i, j, grid)
			if grid[i][j] > val {
				return grid[i][j]
			}
		}
		// Move right currSize times
		for range currSize {
			j++
			fillCell(i, j, grid)
			if grid[i][j] > val {
				return grid[i][j]
			}
		}
	}
	panic("Cannot find")
}

func main() {
	grid := make([][]int, N)
	for i := range N {
		grid[i] = make([]int, N)
	}
	val := 325489
	ans := getFirstBiggerVal(val, grid)
	fmt.Printf("Answer: %d\n", ans)

}
