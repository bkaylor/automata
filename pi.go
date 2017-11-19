package main

import (
	"fmt"
	"math"
	"time"
)

var checked int
var passed int

// Swap 2 slice elements
func swap(x int, y int, digits []int) {
	temp := digits[y]
	digits[y] = digits[x]
	digits[x] = temp
}

// Check if integer m made from first n digits of the number is divisble by n
func check(digits []int) {
	var num int
	var pass bool
	for i := 0; i < len(digits); i++ {
		pass = true
		num = 0	
		temp := digits[0:len(digits)-i]

		for j := 0; j < len(temp); j++ {
			num += int(math.Pow(float64(10), float64(j))) * temp[j]
		}
		if num % len(temp) != 0 {
			pass = false
			break
		}
	}
	if pass == true {
		passed++
		num = 0
		for i := 0; i < len(digits); i++ {
			num += int(math.Pow(float64(10), float64(i))) * digits[i]
		}
		fmt.Println(num)
	}
}

// Heap's Algorithm
func permute(n int, digits []int) {
	if n == 1 {
		checked++
		check(digits)
	} else {
		for i := 0; i < n - 1; i++ {
			permute(n-1, digits)
			if n % 2 == 0 {
				swap(i, n-1, digits)
			} else {
				swap(0, n-1, digits)
			}
		}
		permute(n-1, digits)
	}
}

func main () {
	// Iterate through all 10 digit numbers with distinct digits
	fmt.Println("Pi")
	checked = 0
	passed = 0
	digits := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	
	start := time.Now()
	permute(len(digits), digits)
	elapsed := time.Since(start)
	
	fmt.Println("**********************")
	fmt.Printf("Checked: %d\nPassed: %d\nTime: %s\n", checked, passed, elapsed)
	fmt.Println("**********************")
}
