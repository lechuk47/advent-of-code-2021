package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func main() {
	data, _ := input.ReadFile("input.txt")

	MAX := 2000

	m := make([][]string, MAX)
	for i := 0; i < MAX; i++ {
		m[i] = make([]string, MAX)
		for j := 0; j < MAX; j++ {
			m[i][j] = "."
		}
	}

	xmax := 0
	ymax := 0
	for _, line := range data {
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			_x, _ := strconv.ParseInt(parts[0], 10, 0)
			_y, _ := strconv.ParseInt(parts[1], 10, 0)
			x := int(_x)
			y := int(_y)
			if x > xmax {
				xmax = x
			}
			if y > ymax {
				ymax = y
			}
			m[y][x] = "#"
		}

		if line == "" {
			xmax++
			ymax++
			for y := 0; y < ymax; y++ {
				fmt.Println(m[y][:xmax])
			}
		}

		//fmt.Println("XMAX ", xmax)
		//fmt.Println("YMAX ", ymax)

		if len(parts) == 1 && line != "" {
			parts = strings.Split(line[11:], "=")
			_value, _ := strconv.ParseInt(parts[1], 10, 0)
			value := int(_value)
			if parts[0] == "y" {
				fmt.Println("Fold UP y=", value)
				for y := value + 1; y < ymax; y++ {
					y1 := y
					y0 := value + (value - y)
					for j := 0; j < xmax; j++ {
						if m[y0][j] == "#" || m[y1][j] == "#" {
							m[y0][j] = "#"
						} else {
							m[y0][j] = "."
						}
					}
				}
				ymax = value
				/*				for y := 0; y < ymax; y++ {
								fmt.Println(m[y][:xmax])
							}*/
			} else if parts[0] == "x" {
				fmt.Println("Fold LEFT x=", value)
				for x := value + 1; x < xmax; x++ {
					x1 := x
					x0 := value + (value - x)
					for i := 0; i < ymax; i++ {
						if m[i][x0] == "#" || m[i][x1] == "#" {
							m[i][x0] = "#"
						} else {
							m[i][x0] = "."
						}
					}
				}
				xmax = value
			}
		}

	}

	for y := 0; y < ymax; y++ {
		fmt.Println(m[y][:xmax])
	}

	c := 0
	fmt.Println("XMAX, ", xmax)
	fmt.Println("YMAX, ", ymax)
	for x := 0; x < xmax; x++ {
		for y := 0; y < ymax; y++ {
			if m[y][x] == "#" {
				c++
			}
		}
	}
	fmt.Println("Count: ", c)
}
