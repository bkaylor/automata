// Any live cell with fewer than two live neighbours dies, as if caused by under-population.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by over-population.
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Grid_X = 20
	Grid_Y = 20
	Rand_Count = 40
	Live_Char = "*"
	Dead_Char = " "
	Delay_MS = 300
)

func count_live_neighbors(grid [Grid_X][Grid_Y]string, x int, y int) int {
	liveNeighbors := 0
	if  x > 0 && y > 0 && grid[x-1][y-1] == Live_Char {
		liveNeighbors++
	}
	if  x > 0 && y < Grid_Y-1 && grid[x-1][y+1] == Live_Char {
		liveNeighbors++
	}
	if x < Grid_X-1 && y < Grid_Y-1 && grid[x+1][y+1] == Live_Char {
		liveNeighbors++
	}
	if x < Grid_X-1 && grid[x+1][y] == Live_Char {
		liveNeighbors++
	}
	if x < Grid_X-1 && y > 0 && grid[x+1][y-1] == Live_Char {
		liveNeighbors++
	}
	if y < Grid_Y-1 && grid[x][y+1] == Live_Char {
		liveNeighbors++
	}
	if y > 0 && grid[x][y-1] == Live_Char {
		liveNeighbors++
	}
	if x > 0 && grid[x-1][y] == Live_Char {
		liveNeighbors++
	}
	
	return liveNeighbors
}

func print_grid (grid [Grid_X][Grid_Y]string) {
	for i := 0; i < Grid_X; i++ {
		for j := 0; j < Grid_Y; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func next_generation(grid [Grid_X][Grid_Y]string) ([Grid_X][Grid_Y]string, int) {
	newGrid := grid
	lifeCount := 0
	for i := 0; i < Grid_X; i++ {
		for j := 0; j < Grid_Y; j++ {
			liveNeighbors := count_live_neighbors(grid, i, j)
			if grid[i][j] == Live_Char {
				// Handle live cells
				lifeCount++
				switch {
					case liveNeighbors < 2:
						newGrid[i][j] = Dead_Char
					case liveNeighbors == 2 || liveNeighbors == 3:
						newGrid[i][j] = Live_Char
					case liveNeighbors > 3:
						newGrid[i][j] = Dead_Char
				}
			} else {
				// Handle dead cells
				if liveNeighbors == 3 {
					newGrid[i][j] = Live_Char
				}
			}
		}
	}
	
	return newGrid, lifeCount
}

func main() {
	fmt.Println("Conway's Game of Life")
	var y [Grid_X][Grid_Y]string
	
	// Zero grid
	for i := 0; i < Grid_X; i++ {
		for j := 0; j < Grid_Y; j++ {
			y[i][j] = Dead_Char
		}
	}
	
	// Initialize grid with randoms
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Rand_Count; {
		rx := rand.Intn(Grid_X)
		ry := rand.Intn(Grid_Y)
		if y[rx][ry] == Dead_Char {
			y[rx][ry] = Live_Char
			i++
		}
	}
	
	// Generation
	genCount := 0
	life := Rand_Count
	for {
		fmt.Printf("\nGeneration %d (Life=%d):\n", genCount, life)
		print_grid(y);
		time.Sleep(time.Millisecond * Delay_MS)
		y, life = next_generation(y);
		if life == 0 {
			break
		}
		genCount++
	}
	
	fmt.Printf("Civilization lasted %d generations.\n", genCount)
}