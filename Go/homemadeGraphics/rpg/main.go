package main

import (
	"fmt"
)

type Player struct {
	x    int
	y    int
	char string
}

func main() {

	var graph [25][75]string
	player := Player{1, 1, "\033[32m█\033[0m"}

	CreateFile("save", "testing")
	fmt.Println(ReadFile("save"))

	for {
		for iy, y := range graph {

			for ix, x := range y {
				if iy == player.y && ix == player.x || iy == player.y && ix == player.x+1 {
					x = player.char
				} else {
					x = "\033[34m█\033[0m"
				}
				fmt.Printf(x)
			}
			fmt.Printf("\n")
		}

		var input string
		fmt.Scanf("%v", &input)
		switch input {
		case "x+":
			player.x += 1
		case "x-":
			player.x -= 1
		case "y+":
			player.y += 1
		case "y-":
			player.y -= 1
		}

		fmt.Printf("\033[2J\033[H")
	}
}
