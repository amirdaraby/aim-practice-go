package main

import (
	"math/rand"
	"strconv"

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

var Score int = 0
var ScoreColor rl.Color

func main() {

	rl.InitWindow(screenWidth, screenHeight, "SHOOT CIRCLE")

	rl.SetTargetFPS(30)

	defer rl.CloseWindow()

	ball := Ball{}

	initBall(&ball)

	var life int = 3
	var lifeColor rl.Color
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawCircle(int32(ball.x), int32(ball.y), ball.radius, rl.Maroon)

		if life < 0 {
			rl.EndDrawing()
			rl.SetTargetFPS(1)
			for !rl.WindowShouldClose() {
				rl.BeginDrawing()
				rl.DrawText("GAME OVER, SCORE: "+strconv.FormatInt(int64(Score), 10), 100, 200, int32(32), rl.RayWhite)
				rl.EndDrawing()
			}
			break
		}

		if Score == 0 {
			ScoreColor = rl.White
		} else if Score > 0 {
			ScoreColor = rl.Green
		} else {
			ScoreColor = rl.Red
		}

		if life >= 2 {
			lifeColor = rl.Green
		} else if life < 2 {
			lifeColor = rl.Red
		}
		if life == 0 {
			lifeColor = rl.Maroon
		}

		rl.DrawText("Score:"+strconv.FormatInt(int64(Score), 10), 40, 40, 12, ScoreColor)
		rl.DrawText("Life:"+strconv.FormatInt(int64(life), 10), 150, 40, 12, lifeColor)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), rl.Vector2{X: float32(ball.x), Y: float32(ball.y)}, ball.radius) {
				Score++
			} else {
				life--
				Score--
			}

			initBall(&ball)
		}
		rl.EndDrawing()
	}
}
func initBall(b *Ball) {
	*b = Ball{x: rand.Intn(maxX-minX) + minX, y: rand.Intn(maxY-minY) + minY, radius: 30}
}
