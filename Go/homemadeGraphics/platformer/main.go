package main

import (
	"fmt"
	"time"
)

type Player struct {
	x    int
	y    int
	char string
}

func main() {

	var graph [25][75]string
	var hrz = time.Second / 60
	player := Player{1, 1, "\033[32m█\033[0m"}

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
		time.Sleep(hrz)
		fmt.Printf("\033[2J\033[H")
	}

}
