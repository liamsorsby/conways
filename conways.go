package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
}

func generateGrid() [6][6]bool {
	a := [6][6]bool{
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
	}

	return a
}

func randomiseGrid() [6][6]bool {
	grid := generateGrid()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle( len(grid), func(i, j int) {
		grid[i][j] = grid[j][i]
	})
	return grid
}

func main() {
	print("\033[H\033[2J")

	grid := randomiseGrid()

	for {
		for _, items := range grid {
			for _, item := range items {
				if item {
					fmt.Print(" X ")
				} else {
					fmt.Print(" 0 ")
				}
			}
			fmt.Println()
		}
		time.Sleep(1 * time.Second)

		print("\033[H\033[2J")
	}
}
