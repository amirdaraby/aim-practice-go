package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	x, y   int
	radius float32
}

const screenHeight int32 = 450
const screenWidth int32 = 800
const minX, maxX int = 70, 750
const minY, maxY int = 70, 400

func main() {

	rl.InitWindow(screenWidth, screenHeight, "SHOOT CIRCLE")

	defer println("shutting down")
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	ball := Ball{}

	initBall(&ball)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.DrawCircle(int32(ball.x), int32(ball.y), ball.radius, rl.Maroon)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			fmt.Println(&ball)
			initBall(&ball)
		}
		rl.EndDrawing()
	}
}

func initBall(b *Ball) {
	*b = Ball{x: rand.Intn(maxX-minX) + minX, y: rand.Intn(maxY-minY) + minY, radius: 30}
}
