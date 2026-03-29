package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const sizex = 500
const sizey = 300

type color struct {
	hue float32
	sat float32
	val float32
}

func main() {
	rl.InitWindow(sizex, sizey, "game")
	rl.SetTargetFPS(60)

	var x, y int32 = 10, 10
	rainbow := color{0, 1, 1}
	direction := rl.Vector2{2, 2}
	for !rl.WindowShouldClose() {
		if rainbow.hue == 360 {
			rainbow.hue = 0
		} else {
			rainbow.hue += 1
		}

		x += int32(direction.X)
		y += int32(direction.Y)
		if x > sizex-65 || x < 0 {
			direction.X *= -1
		}
		if y > sizey-25 || y < 0 {
			direction.Y *= -1
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawText("DVD", x, y, 20, rl.ColorFromHSV(rainbow.hue, rainbow.sat, rainbow.val))

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
