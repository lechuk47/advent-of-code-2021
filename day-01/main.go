package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInput(path string) ([]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	measurements := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m, _ := strconv.ParseInt(line, 10, 0)
		measurements = append(measurements, m)
	}
	return measurements, scanner.Err()
}

func CountIncrements(input []int64) int {
	increments := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increments++
		}
	}
	return increments
}

func CountIncrementsExtended(input []int64) int {
	increments := 0
	for i := 0; i < len(input)-3; i++ {
		a := input[i]
		b := input[i+3]
		// No need to sum all values, input[i+1] and input[i+2] are common
		// a := input[i] + input[i+1] + input[i+2]
		// b := input[i+1] + input[i+2] + input[i+3]
		if b > a {
			increments++
		}

	}
	return increments
}

func main() {
	//input = []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	input, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	increments := CountIncrements(input)
	fmt.Printf("Increments: %d\n", increments)

	increments = CountIncrementsExtended(input)
	fmt.Printf("Increments: %d\n", increments)

}
