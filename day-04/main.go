package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

type board struct {
	Rows [][]int
	Cols [][]int
}

func NewBoard() board {
	b := board{}
	b.Rows = make([][]int, 5)
	b.Cols = make([][]int, 5)
	for i := 0; i < 5; i++ {
		b.Rows[i] = make([]int, 5)
		b.Cols[i] = make([]int, 5)
	}
	return b
}

func (b *board) Init() {
	// Traspose Matrix
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b.Cols[i][j] = b.Rows[j][i]
		}
	}
	// Sort Matrixes
	for i := 0; i < 5; i++ {
		sort.Ints(b.Rows[i])
		sort.Ints(b.Cols[i])
	}
}

func (b *board) MarkNumber(n int) {
	fmt.Println("MarkNumber")
}
func (b *board) CheckLines(numbers []int) int {

	minR := 101
	minC := 101

	for i := 0; i < 5; i++ {
		// Check all rows and columns
		matchR := 0
		matchC := 0

		//for n := 0; n < len(numbers); n++ {
		for n, number := range numbers {
			for j := 0; j < 5; j++ {

				// Checking Row[i]
				if b.Rows[i][j] == number {
					matchR++
					if matchR == 5 {
						if minR == -1 || n < minR {
							minR = n
						}

					}
				}
				// Checking Col[i]
				if b.Cols[i][j] == number {
					matchC++
					if matchC == 5 {
						if minC == -1 || n < minC {
							minC = n
						}
					}
				}
			}
		}
	}
	min := minR
	if minC < minR {
		min = minC
	}
	return min
}

func InSlice(numbers []int, number int) bool {
	for _, v := range numbers {
		if v == number {
			return true
		}
	}
	return false
}

func (b *board) GetResult(numbers []int, last int) int {
	//sort.Ints(numbers)
	result := 0
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {

			if !InSlice(numbers, b.Rows[i][j]) {
				result += b.Rows[i][j]
			}
			if i != j {
				if !InSlice(numbers, b.Rows[j][i]) {
					result += b.Rows[j][i]
				}
			}
		}
	}
	return result * last
}

func main() {
	input, err := input.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var boards []board
	row := 0
	var inputRow []string
	b := NewBoard()

	for i := 2; i < len(input); i++ {
		if len(input[i]) > 0 {
			// Line is not empty
			inputRow = strings.Fields(input[i])
			for j := 0; j < len(inputRow); j++ {
				n, _ := strconv.ParseInt(inputRow[j], 10, 0)
				b.Rows[row][j] = int(n)
			}
			row++
		}
		if row > 4 {
			b.Init()
			boards = append(boards, b)
			b = NewBoard()
			row = 0
		}
	}

	inputList := strings.Split(input[0], ",")
	numbers := make([]int, len(inputList))
	for i := 0; i < len(inputList); i++ {
		n, _ := strconv.ParseInt(inputList[i], 10, 0)
		numbers[i] = int(n)
	}

	var winner board
	var lastWinner board
	minIndex := -1
	maxIndex := -1

	for b := 0; b < len(boards); b++ {
		min := boards[b].CheckLines(numbers)
		if minIndex < 0 || min < minIndex {
			winner = boards[b]
			minIndex = min

		}
		if maxIndex < 0 || min > maxIndex {
			lastWinner = boards[b]
			maxIndex = min
		}
	}
	lastW := numbers[minIndex]
	lastL := numbers[maxIndex]

	resultW := winner.GetResult(numbers[0:minIndex+1], lastW)
	fmt.Println("Winner:", resultW)
	fmt.Println("LastW:", lastW)
	fmt.Println("MinIndex:", minIndex)

	resultL := lastWinner.GetResult(numbers[0:maxIndex+1], lastL)
	fmt.Println("lastWinner:", resultL)
	fmt.Println("LastL:", lastL)
	fmt.Println("MaxIndex:", maxIndex)

}
