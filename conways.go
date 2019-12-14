package main

import (
	"flag"
	"fmt"
	"time"

	"stash.skybet.net/~sorsbyl/conways/templates"
	"stash.skybet.net/~sorsbyl/conways/utils"
)

var template string

func init() {
	flag.StringVar(&template, "", "", "Choose the template that is going to render on screen")
}

func main() {
	flag.Parse()
	clearScreen()
	grid := templates.New(template)
	printGrid(grid)
}

func printGrid(grid [][]bool) {
	for {
		for _, row := range grid {
			for _, item := range row {
				if item {
					fmt.Print("⬜")
				} else {
					fmt.Print("⬛")
				}
			}
			fmt.Println()
		}

		time.Sleep(700 * time.Millisecond)
		clearScreen()
		grid = utils.Evaluate(grid)
	}
}

func clearScreen() {
	print("\033[H\033[2J")
}
