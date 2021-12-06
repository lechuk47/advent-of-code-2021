package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func IndexExists(elements []int, element int) bool {
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return true
		}
	}
	return false
}

func RemoveIndexes(s []string, indexes *[]int) []string {
	new_input := []string{}
	for i := 0; i < len(s); i++ {
		if !IndexExists(*indexes, i) {
			new_input = append(new_input, s[i])
		}
	}
	return new_input
}

func getBitIndexes(input []string, position int) ([]int, []int) {
	index1 := []int{}
	index0 := []int{}

	for i, e := range input {
		if string(e[position]) == "1" {
			index1 = append(index1, i)
		} else {
			index0 = append(index0, i)
		}
	}
	return index0, index1
}

func GetNewInput(input []string, position int, preference int, binsize int) []string {
	index1, index0 := getBitIndexes(input, position)
	var to_remove *[]int

	to_remove = &index1
	if preference == 1 {
		if len(index1) > len(index0) {
			to_remove = &index0
		}
	} else {
		if len(index1) <= len(index0) {
			to_remove = &index0
		}
	}
	new_input := RemoveIndexes(input, to_remove)
	return new_input
}

func Part01(input []string) {

	binsize := len(input[0])
	input_size := len(input)
	ones := make([]int, binsize)

	for _, line := range input {
		for i, c := range line {
			if string(c) == "1" {
				ones[i] += 1
			}
		}
	}

	epsilon := make([]string, binsize)
	gamma := make([]string, binsize)

	for i, n := range ones {
		if n > (input_size / 2) {
			epsilon[i] = "1"
			gamma[i] = "0"
		} else {
			epsilon[i] = "0"
			gamma[i] = "1"
		}

	}
	e := strings.Join(epsilon, "")
	ee, _ := strconv.ParseInt(e, 2, 0)
	g := strings.Join(gamma, "")
	gg, _ := strconv.ParseInt(g, 2, 0)
	fmt.Println("Epsilon: ", e)
	fmt.Println("Gamma: ", gg)
	fmt.Println("Result: ", gg*ee)
}

func Part02(input []string) {
	binsize := len(input[0])
	o2 := input
	co2 := input
	for p := 0; p < binsize; p++ {
		if len(o2) > 1 {
			o2 = GetNewInput(o2, p, 1, binsize)
		}
		if len(co2) > 1 {
			co2 = GetNewInput(co2, p, 0, binsize)

		}
		if len(o2) <= 1 && len(co2) <= 1 {
			break
		}
	}

	oo2, _ := strconv.ParseInt(o2[0], 2, 0)
	cco2, _ := strconv.ParseInt(co2[0], 2, 0)
	fmt.Println("O2: ", oo2)
	fmt.Println("Co2: ", cco2)
	fmt.Println("Result: ", oo2*cco2)
}

func main() {
	//input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	input, err := input.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Part01(input)
	Part02(input)

}
