package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

func PrintMatrix(m [][]int) {

	for _, l := range m {
		fmt.Println(l)
	}
}

func FindLowPoints(m [][]int) ([]int, []int) {
	var lowPoints []int
	var bassinSizes []int
	for i := 0; i < len(m); i++ {
		up, down := true, true
		if i == 0 {
			up = false
		}
		if i == len(m)-1 {
			down = false
		}
		for j := 0; j < len(m[i]); j++ {
			left, right := true, true
			if j == 0 {
				left = false
			}
			if j == len(m[i])-1 {
				right = false
			}
			n := m[i][j]

			if up && m[i-1][j] <= n {
				continue
			}

			if down && m[i+1][j] <= n {
				continue
			}

			if left && m[i][j-1] <= n {
				continue
			}

			if right && m[i][j+1] <= n {
				continue
			}

			lowPoints = append(lowPoints, n)

			// BASSINS
			bSize := findBassinSize(m, i, j)
			bassinSizes = append(bassinSizes, bSize)
		}
	}
	return lowPoints, bassinSizes
}

type point struct {
	i, j int
}

func findPoint(points []point, p point) bool {
	for _, v := range points {
		if p.i == v.i && p.j == v.j {
			return true
		}
	}
	return false
}

func findBassinSize(m [][]int, ii, jj int) int {
	var managedPoints []point
	var queuePoints []point
	//
	size := 0
	s := point{i: ii, j: jj}
	queuePoints = append(queuePoints, s)

	for len(queuePoints) > 0 {
		// Pop element
		p := queuePoints[0]
		queuePoints = queuePoints[1:len(queuePoints)]
		if !findPoint(managedPoints, p) {
			// Count element
			size++
			managedPoints = append(managedPoints, p)
		}
		// up
		i := p.i
		j := p.j
		for i > 0 && m[i-1][j] != 9 && m[i-1][j] > m[i][j] {
			np := point{i: i - 1, j: j}
			if !findPoint(managedPoints, np) {
				queuePoints = append(queuePoints, np)
			}
			i--
		}
		// Down
		i = p.i
		j = p.j
		for i < len(m)-1 && m[i+1][j] != 9 && m[i+1][j] > m[i][j] {
			np := point{i: i + 1, j: j}
			if !findPoint(managedPoints, np) {
				queuePoints = append(queuePoints, np)
			}
			i++
		}
		// Find mostleft
		i = p.i
		j = p.j
		for j > 0 && m[i][j-1] != 9 && m[i][j-1] > m[i][j] {
			np := point{i: i, j: j - 1}
			if !findPoint(managedPoints, np) {
				queuePoints = append(queuePoints, np)
			}
			j--
		}
		// Find mostRight
		i = p.i
		j = p.j
		for j < len(m[i])-1 && m[i][j+1] != 9 && m[i][j+1] > m[i][j] {
			np := point{i: i, j: j + 1}
			if !findPoint(managedPoints, np) {
				queuePoints = append(queuePoints, np)
			}
			j++
		}
	}

	return size

}

func main() {
	data, _ := input.ReadFile("input.txt")

	matrix := make([][]int, len(data))
	for i, line := range data {
		matrix[i] = make([]int, len(line))
		v := strings.Split(line, "")
		for j, k := range v {
			n, _ := strconv.ParseInt(k, 10, 0)
			matrix[i][j] = int(n)
		}
	}
	//PrintMatrix(matrix)
	lp, bs := FindLowPoints(matrix)
	totalRisk := 0
	for _, n := range lp {
		totalRisk += n + 1
	}
	//fmt.Println("LowPoints: ", lp)
	fmt.Println("TotalRisk: ", totalRisk)
	//fmt.Println("BasinSizes: ", bs)
	sort.Ints(bs)
	n := len(bs)
	fmt.Println("3LargestBassinsSize: ", bs[n-1]*bs[n-2]*bs[n-3])

}
