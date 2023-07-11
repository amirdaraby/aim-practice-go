package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(800, 450, "this is a test")
	defer rl.CloseWindow()
	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("This is a test, thanks for watching", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
