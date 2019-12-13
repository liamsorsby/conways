package main

import (
	"fmt"
	"time"
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
