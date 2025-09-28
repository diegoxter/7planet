package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/diegoxter/7planet/internal/game"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(600)

	rl.InitWindow(screenWidth, screenHeight, "7th Planet (alpha)")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	g := game.Init(screenWidth, screenHeight)
	defer g.Unload()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(*g.Camera)
		g.Run()
		rl.EndMode2D()

		rl.DrawText("this IS a texture loaded from an image!", 300, 370, 10, rl.Gray)

		rl.EndDrawing()
	}
}
