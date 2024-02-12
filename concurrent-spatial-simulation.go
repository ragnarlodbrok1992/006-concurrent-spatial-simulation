package main

import (
  "fmt"
  raylib "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH int32 = 1280
const WINDOW_HEIGHT int32 = 720
const FONT_NAME string = "assets/fonts/SourceCodePro-Medium.ttf"
const DEFAULT_FONT_SIZE float32 = 20
const DEFAULT_SPACING float32 = 2

func drawTextDefault(font raylib.Font, text string, x float32, y float32) {
  raylib.DrawTextEx(font, text, raylib.Vector2{x, y}, DEFAULT_FONT_SIZE, DEFAULT_SPACING, raylib.DarkGray)
}

func main() {
  // Start message
  fmt.Println("Concurrent Spatial Simulation System - early alpha version.")
  defer fmt.Println("All went well, goodbye...")

  // Initialization
  raylib.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Concurrent Spatial Simulation System")
  defer raylib.CloseWindow()

  // Configuration
  raylib.SetTargetFPS(60)

  // Loading different font
  DEFAULT_FONT := raylib.LoadFont(FONT_NAME)

  for !raylib.WindowShouldClose() {
    // Input handling
    if raylib.IsKeyPressed(raylib.KeyEscape) {
      raylib.CloseWindow()
    }

    // Simulation

    // Frame drawing
    raylib.BeginDrawing()
    raylib.ClearBackground(raylib.RayWhite)

    drawTextDefault(DEFAULT_FONT, "Concurrent Spatial Simulation System", 10, 10)

    raylib.EndDrawing()
  }

}
