package main

import (
	"fmt"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func main() {
	data, _ := input.ReadFile("input.txt")
	craps, _ := input.ConvertStringToSliceOfInts(data[0], ",")
	//craps := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	crapsPosition := make(map[int]int)

	// CLassify craps per position
	var maxPosition int
	for _, c := range craps {
		crapsPosition[c]++
		if c > maxPosition {
			maxPosition = c
		}
	}
	// Generate a Slice with the cost value per number of moves
	moveCosts := make([]int, maxPosition+1)
	moveCosts[0] = 0
	for i := 1; i <= maxPosition; i++ {
		moveCosts[i] = (i * i) - moveCosts[i-1]
	}

	min := -1
	minPos := -1
	var fuel int
	var moves int
	//for p, _ := range crapsPosition {
	for p := 0; p < maxPosition; p++ {
		fuel = 0
		for k, v := range crapsPosition {
			if k == p {
				continue
			}
			moves = k - p
			if moves < 0 {
				moves = moves * -1
			}
			fuel += moveCosts[moves] * v
		}
		if min == -1 || fuel < min {
			min = fuel
			minPos = p
		}
	}
	fmt.Println("Fuel: ", min)
	fmt.Println("MinPos: ", minPos)
}
