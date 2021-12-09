package main

import (
	"fmt"
	"os"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func GetFishInput() []int {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fish, err := input.ConvertStringToSliceOfInts(data[0], ",")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fish
}

func main() {
	//fish := []int{3, 4, 3, 1, 2}
	fish := GetFishInput()
	days := 256

	// count fishes by its internal timer
	fishTimers := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	// Classify all the fishes
	for _, f := range fish {
		fishTimers[f]++
	}

	for d := 0; d < days; d++ {
		new := fishTimers[0]
		fishTimers = fishTimers[1:]
		fishTimers = append(fishTimers, 0)
		fishTimers[6] += new
		fishTimers[8] += new
	}

	nfish := 0
	for _, n := range fishTimers {
		nfish += n
	}
	fmt.Println("Fish: ", nfish)
	fmt.Println("FishTimers: ", fishTimers)
}
