package main

import (
  "fmt"
  grid "github.com/ragnarlodbrok1992/concurrent-spatial-simulation/src/grid"
  raylib "github.com/gen2brain/raylib-go/raylib"
)

const WINDOW_WIDTH int32 = 1280
const WINDOW_HEIGHT int32 = 720
const FONT_NAME string = "assets/fonts/SourceCodePro-Medium.ttf"
const DEFAULT_FONT_SIZE float32 = 20
const DEFAULT_SPACING float32 = 2

func drawTextDefault(font raylib.Font, text string, x float32, y float32) {
  raylib.DrawTextEx(font, text, raylib.Vector2{x, y}, DEFAULT_FONT_SIZE, DEFAULT_SPACING, raylib.Black)
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
  BACKGROUND_MAIN_COLOR := raylib.Color{252, 115, 3, 255}

  // Debug code
  // grid.DrawGrid()
  debug_grid := grid.CreateGrid(10, 10)
  fmt.Println(debug_grid)

  for !raylib.WindowShouldClose() {
    // Input handling
    if raylib.IsKeyPressed(raylib.KeyEscape) {
      raylib.CloseWindow()
    }

    // Simulation

    // Frame drawing
    raylib.BeginDrawing()
    raylib.ClearBackground(BACKGROUND_MAIN_COLOR)

    drawTextDefault(DEFAULT_FONT, "Concurrent Spatial Simulation System", 10, 10)

    raylib.EndDrawing()
  }

}
