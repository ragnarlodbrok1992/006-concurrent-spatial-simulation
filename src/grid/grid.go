package grid

import "fmt"

// Grid metadata structure definition

// Grid structure definition
// TODO: Think about container for it?
//       Grids will be mostly static but who knows.

// GridCell for game of life

type CellState int

const (
  ALIVE CellState = iota
  DEAD
)

type GridCell struct {
  x, y int
  id int
  state CellState
}

func DrawGrid() {
  fmt.Println("Drawing grid!")
}

func SimulateGrid() { // Those go in main loop so no printing here
  // fmt.Println("Simulating grid!")
}

func CreateGrid(num_x int, num_y int) []GridCell { // Those go in main loop so no printing here
  local_id := 0
  GridSlice := make([]GridCell, 0, num_x * num_y)
  for i := 0; i < num_x; i++ {
    for j := 0; j < num_y; j++ {
      GridSlice = append(GridSlice, GridCell{x: i, y: j, id: local_id, state: DEAD})
      local_id++
    }
  }

  return GridSlice
  // fmt.Println("Creating grid!")
}
