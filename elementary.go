package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

const (
	Generations = 20
	Live_Char = "*"
	Dead_Char = " "
	Delay_MS = 30
)

func next_generation (oldField []string, Rule uint8) []string {
	Width := (Generations*2) + 1
	newFieldRaw := make ([]uint8, Width, Width)
	newField := make ([]string, Width, Width)
	var neighbors [3]int
	
	// Implement
	
	for i := 0; i < Width; i++ {
		neighbors = [3]int{0, 0, 0}
		if i != 0 && oldField[i-1] == Live_Char {
			neighbors[0] = 1
		}
		if oldField[i] == Live_Char {
			neighbors[1] = 1
		}
		if i < Width-1 && oldField[i+1] == Live_Char {
			neighbors[2] = 1
		}
		
		switch neighbors {
			case [3]int{0, 0, 0}: newFieldRaw[i] = (1 << 0) & Rule
			case [3]int{0, 0, 1}: newFieldRaw[i] = (1 << 1) & Rule
			case [3]int{0, 1, 0}: newFieldRaw[i] = (1 << 2) & Rule
			case [3]int{0, 1, 1}: newFieldRaw[i] = (1 << 3) & Rule
			case [3]int{1, 0, 0}: newFieldRaw[i] = (1 << 4) & Rule
			case [3]int{1, 0, 1}: newFieldRaw[i] = (1 << 5) & Rule
			case [3]int{1, 1, 0}: newFieldRaw[i] = (1 << 6) & Rule
			case [3]int{1, 1, 1}: newFieldRaw[i] = (1 << 7) & Rule
		}
	}
	
	for i := 0; i < Width; i++ {
		if newFieldRaw[i] != 0 {
			newField[i] = Live_Char
		} else {
			newField[i] = Dead_Char
		}
	}
	
	return newField
}

func print_field(field []string) {
	for i := 0; i < len(field); i++ {
		fmt.Printf("%s ", field[i])
	}
	fmt.Println()
}

func main() {
	fmt.Println("Elementary Cellular Automata")
	
	var Rule64 uint64

	if len(os.Args) > 1 {
		Rule64, _ = strconv.ParseUint(os.Args[1], 10, 8)
	} else {
		fmt.Println("Usage: automata [rule number]")
		Rule64 = 30
	}
	
	Rule := uint8(Rule64)
	
	// Compute field width
	Width := (Generations*2) + 1
	field := make([]string, Width, Width)
	
	// Zero field
	for i := 0; i < Width; i++ {
		field[i] = Dead_Char
	}
	
	// Starting cell
	field[Width/2] = Live_Char
	
	for currentGen := 0; currentGen < Generations; currentGen++ {
		// fmt.Printf("Generation %d:\n", currentGen)
		print_field(field)
		field = next_generation(field, Rule)
		time.Sleep(time.Millisecond * Delay_MS)
	}
}

