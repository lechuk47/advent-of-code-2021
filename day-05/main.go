package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

type Point struct {
	x, y int
}

type Vector struct {
	a Point
	b Point
}
type Map struct {
	matrix  [][]int
	vectors []Vector
}

const MAX = 1000

func NewMap() Map {
	m := Map{}
	m.matrix = make([][]int, MAX)
	for i := 0; i < MAX; i++ {
		m.matrix[i] = make([]int, MAX)
	}
	return m
}

func Sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func (m *Map) GenerateMatrix() {

	for _, v := range m.vectors {
		ax := v.a.x
		ay := v.a.y
		bx := v.b.x
		by := v.b.y
		//fmt.Println("L ", x1, y1, "---", x2, y2)
		if ax == bx {
			// Vertical
			x := ax
			y, yy := Sort(ay, by)
			//fmt.Println("H ", x1, y1, x2, y2)
			for j := y; j <= yy; j++ {
				m.matrix[j][x]++
			}
		} else if ay == by {
			// Horizontal
			//fmt.Println("V ", x1, y1, x2, y2)
			y := ay
			x, xx := Sort(ax, bx)
			for i := x; i <= xx; i++ {
				m.matrix[y][i]++
			}
		} else {
			// DIAGONAL
			// Mange diagonals to the right only
			// if the line goes to the right, switch points
			// s -> Start Point
			// e -> End Point
			var sx, sy, ex, ey int
			if bx > ax { // -->
				sx = ax
				sy = ay
				ex = bx
				ey = by
			} else { // <--
				sx = bx
				sy = by
				ex = ax
				ey = ay
			}
			if ey > sy {
				// 1,1 - 2,2 - 3,3
				// 2,0 - 3,1 -- 4,2 -- 5,3 -- 6,4
				for {
					if sx <= ex && sy <= ey {
						m.matrix[sy][sx]++
						sx++
						sy++
					} else {
						break
					}
				}
			} else {
				// 7,9 -- 8,8 -- 9,7
				// 2,2 -- 3,1 -- 4,0
				for {
					if sx <= ex && sy >= ey {
						m.matrix[sy][sx]++
						sx++
						sy--
					} else {
						break
					}
				}
			}
		}
	}
}

func (m Map) PrintMatrix() {
	for _, x := range m.matrix[0:MAX] {
		fmt.Println(x[0:MAX])
	}
}

func (m Map) CountOverlapping() int {
	count := 0
	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			if m.matrix[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	input, err := input.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := NewMap()

	for _, line := range input {
		elements := strings.Split(line, " -> ")
		_a := strings.Split(elements[0], ",")
		_b := strings.Split(elements[1], ",")
		xa, _ := strconv.ParseInt(_a[0], 10, 0)
		ya, _ := strconv.ParseInt(_a[1], 10, 0)
		xb, _ := strconv.ParseInt(_b[0], 10, 0)
		yb, _ := strconv.ParseInt(_b[1], 10, 0)
		a := Point{x: int(xa), y: int(ya)}
		b := Point{x: int(xb), y: int(yb)}
		v := Vector{a: a, b: b}
		m.vectors = append(m.vectors, v)
	}
	m.GenerateMatrix()
	//m.PrintMatrix()
	overlapping := m.CountOverlapping()
	fmt.Println("Overlaps: ", overlapping)
}
