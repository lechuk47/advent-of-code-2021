package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func CountIncrements(input []int) int {
	increments := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increments++
		}
	}
	return increments
}

func CountIncrementsExtended(input []int) int {
	increments := 0
	a := input[0]
	for i := 3; i < len(input)-3; i += 3 {
		b := input[i]
		if b > a {
			increments++
		}
		a = input[i-3]

	}
	return increments
}

func main() {
	//input = []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	input, err := input.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var measurements = []int{}
	for _, line := range input {
		m, _ := strconv.ParseInt(line, 10, 0)
		measurements = append(measurements, int(m))
	}
	increments := CountIncrements(measurements)
	fmt.Printf("Increments: %d\n", increments)

	increments = CountIncrementsExtended(measurements)
	fmt.Printf("Increments: %d\n", increments)

}
