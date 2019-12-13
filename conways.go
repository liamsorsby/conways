package main

import (
	"fmt"
	"time"

	"stash.skybet.net/~sorsbyl/conways/utils"
)

func generateInitialGrid() [][]bool {
	return [][]bool{
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, true, true, true, false},
		{false, true, true, true, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
	}
}

func printGrid(grid [][]bool) {
	for {
		for _, row := range grid {
			for _, item := range row {
				if item {
					fmt.Print(" ⬛ ")
				} else {
					fmt.Print(" ⬜ ")
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
	grid := generateInitialGrid()
	printGrid(grid)
}
