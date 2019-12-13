package main

import (
	"fmt"
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

func main() {
	print("\033[H\033[2J")

	grid := generateGrid()

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
