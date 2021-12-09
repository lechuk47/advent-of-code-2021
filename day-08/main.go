package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func Difference(a string, b string) []string {
	// Returns elements in group A not in B
	aa := strings.Split(a, "")
	bb := strings.Split(b, "")
	var elements []string
	for _, i := range aa {
		found := false
		for _, j := range bb {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			elements = append(elements, i)
		}
	}
	return elements
}

func Intersect(a string, b string) []string {
	aa := strings.Split(a, "")
	sort.Strings(aa)
	bb := strings.Split(b, "")
	sort.Strings(bb)
	var elements []string
	for _, i := range aa {
		for _, j := range bb {
			if i == j {
				elements = append(elements, i)
			}
		}
	}
	return elements
}

func SortCode(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func GetOutputValue(patterns []string, output []string) int {

	numberCode := make([]string, 10)
	var patterns5 []string
	var patterns6 []string
	for _, p := range patterns {
		pos := -1
		switch len(p) {
		case 2:
			pos = 1
		case 3:
			pos = 7
		case 4:
			pos = 4
		case 5:
			patterns5 = append(patterns5, p)
		case 6:
			patterns6 = append(patterns6, p)
		case 7:
			pos = 8
		}
		if pos != -1 {
			numberCode[pos] = SortCode(p)
		}
	}

	//3
	for i, p := range patterns5 {
		t := Intersect(p, numberCode[1])
		if len(t) == 2 {
			numberCode[3] = SortCode(p)
			patterns5 = append(patterns5[:i], patterns5[i+1:]...)
			break
		}
	}

	//9
	for i, p := range patterns6 {
		d := Difference(p, numberCode[3])
		if len(d) == 1 {
			numberCode[9] = SortCode(p)
			patterns6 = append(patterns6[:i], patterns6[i+1:]...)
			break
		}
	}

	//5, 2
	for _, p := range patterns5 {
		d := Difference(numberCode[9], p)
		if len(d) == 1 {
			numberCode[5] = SortCode(p) //strings.Split(p, "")
		}
		if len(d) == 2 {
			numberCode[2] = SortCode(p) // strings.Split(p, "")
		}
	}

	// 0,6
	for _, p := range patterns6 {
		i := Difference(p, numberCode[5])
		if len(i) == 1 {
			numberCode[6] = SortCode(p) // strings.Split(p, "")
		}
		if len(i) == 2 {
			numberCode[0] = SortCode(p) // strings.Split(p, "")
		}
	}

	numberByCode := make(map[string]int, 10)
	for i, n := range numberCode {
		numberByCode[n] = i
	}

	result := 0
	for i, o := range output {
		code := SortCode(o)
		p := float64((i - 3) * -1)
		pow := int(math.Pow(10, p))
		result += numberByCode[code] * pow
	}
	return result
}

func main() {
	data, _ := input.ReadFile("input.txt")
	//craps, _ := input.ConvertStringToSliceOfInts(data[0], ",")
	//craps := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	total := 0
	for _, line := range data {
		values := strings.Split(line, "|")
		o := strings.Fields(values[1])
		p := strings.Fields(values[0])
		ov := GetOutputValue(p, o)
		total += ov
	}
	fmt.Println("Total: ", total)
}
