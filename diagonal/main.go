package main

import "fmt"

func main() {
	grid := [][]int{
		{7, 11},
		{3, 15},
	}
	res := diagonalTraverse(grid)
	for i := range res {
		fmt.Print(i)
	}
}

func diagonalTraverse(grid [][]int) []int {
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	out := []int{}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			out = append(out, grid[x][y])
		}

	}
	return out
}
