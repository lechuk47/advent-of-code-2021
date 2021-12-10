package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func findInSlice(list []string, s string) int {
	for i, v := range list {
		if s == v {
			return i
		}
	}
	return -1
}

func main() {

	openers := []string{"(", "[", "{", "<"}
	closers := []string{")", "]", "}", ">"}
	points := []int{3, 57, 1197, 25137}

	lines, _ := input.ReadFile("input.txt")
	errorScore := 0
	var completionScores []int
	for _, l := range lines {
		stack := []string{}
		lineSimbols := strings.Split(l, "")
		for _, s := range lineSimbols {
			i := findInSlice(openers, s)
			if i != -1 {
				stack = append(stack, s)
			} else {
				j := findInSlice(closers, s)
				if j != -1 {
					lastOpenerIndex := findInSlice(openers, stack[len(stack)-1])
					if j == lastOpenerIndex {
						stack = stack[:len(stack)-1]
						lastOpenerIndex = len(stack) - 1
					} else {
						fmt.Println("Corrupt line: ", l)
						errorScore += points[j]
						break
					}
				}
			}
		}

		if len(stack) > 0 {
			fmt.Println("Incomplete line: ", l)
			lineScore := 0
			for i := len(stack) - 1; i >= 0; i-- {
				k := findInSlice(openers, stack[i])
				lineScore = 5 * lineScore
				lineScore += k + 1
			}
			completionScores = append(completionScores, lineScore)
		}
	}
	sort.Ints(completionScores)
	fmt.Println("SyntaxErrorScore: ", errorScore)
	fmt.Println("CompletionScore: ", completionScores[len(completionScores)/2])
}
