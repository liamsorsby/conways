package main

import (
	"fmt"
	"math/rand"
	"time"

	"stash.skybet.net/~sorsbyl/conways/utils"
)

func generateInitialGrid() [6][6]bool {
	return [6][6]bool{
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
	}
}

func randomiseGrid() [6][6]bool {
	grid := generateInitialGrid()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(grid), func(i, j int) {
		grid[i][j] = grid[j][i]
	})
	return grid
}

func printGrid(grid [6][6]bool) {
	for {
		for _, row := range grid {
			for _, item := range row {
				if item {
					fmt.Print(" X ")
				} else {
					fmt.Print(" 0 ")
				}
			}
			fmt.Println()
		}

		time.Sleep(1 * time.Second)
		clearScreen()
		grid = utils.Evaluate(grid)
	}
}

func clearScreen() {
	print("\033[H\033[2J")
}

func main() {
	clearScreen()
	grid := randomiseGrid()
	printGrid(grid)
}
