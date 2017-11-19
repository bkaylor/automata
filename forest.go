// A burning cell turns into an empty cell
// A tree will burn if at least one neighbor is burning
// A tree ignites with probability f even if no neighbor is burning
// An empty space fills with a tree with probability p

package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	Grid_X = 20
	Grid_Y = 40
	Fire_Prob = 0.001
	Life_Prob = 0.05
	Live_Char = "O"
	Burning_Char = "X"
	Empty_Char = " "
	Delay_MS = 300
)

func burning_neighbors(grid [Grid_X][Grid_Y]int, x int, y int) int {
	burningNeighbors := 0
	if  x > 0 && y > 0 && grid[x-1][y-1] == 2 {
		burningNeighbors++
	}
	if  x > 0 && y < Grid_Y-1 && grid[x-1][y+1] == 2 {
		burningNeighbors++
	}
	if x < Grid_X-1 && y < Grid_Y-1 && grid[x+1][y+1] == 2 {
		burningNeighbors++
	}
	if x < Grid_X-1 && grid[x+1][y] == 2 {
		burningNeighbors++
	}
	if x < Grid_X-1 && y > 0 && grid[x+1][y-1] == 2 {
		burningNeighbors++
	}
	if y < Grid_Y-1 && grid[x][y+1] == 2 {
		burningNeighbors++
	}
	if y > 0 && grid[x][y-1] == 2 {
		burningNeighbors++
	}
	if x > 0 && grid[x-1][y] == 2 {
		burningNeighbors++
	}
	
	return burningNeighbors
}

func print_grid(grid [Grid_X][Grid_Y]int) {
	for i := 0; i < Grid_X; i++ {
		for j := 0; j < Grid_Y; j++ {
			switch grid[i][j]{
				case 0: fmt.Print(Empty_Char)
				case 1: fmt.Print(Live_Char)
				case 2: fmt.Print(Burning_Char)
			}
		}
		fmt.Println()
	}
}

func simstep(grid [Grid_X][Grid_Y]int) [Grid_X][Grid_Y]int {
	var newGrid [Grid_X][Grid_Y]int	
	for i := 0; i < Grid_X; i++ {
		for j := 0; j < Grid_Y; j++ {
			switch grid[i][j] {
				case 0: 
					if rand.Intn(1/Life_Prob) == 0 {
						newGrid[i][j] = 1
					}
					
				case 1: 
					newGrid[i][j] = 1
					if burning_neighbors(grid, i, j) > 0 {
						newGrid[i][j] = 2
					} else {
						if rand.Intn(1/Fire_Prob) == 0 {
							newGrid[i][j] = 2
						}
					}
					
				case 2: 
					newGrid[i][j] = 0
			}
		}
	}

	return newGrid
	
}

func main() {
	fmt.Println("Forest Fire")
	var grid [Grid_X][Grid_Y]int
	rand.Seed(time.Now().UnixNano())
	
	for step := 0;; step++ {
		fmt.Printf("\nStep %d:\n", step)
		print_grid(grid)
		grid = simstep(grid)
		time.Sleep(time.Millisecond * Delay_MS)
	}
}