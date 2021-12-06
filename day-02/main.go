package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/lechuk47/advent-of-code-2021/input"
)

type position struct {
	hpos  int64
	depth int64
	aim   int64
}

func (p *position) Forward(value int64) {
	p.hpos = p.hpos + value
	p.depth = p.depth + (p.aim * value)
}

func (p *position) Down(value int64) {
	//p.depth += value
	p.aim += value
}

func (p *position) Up(value int64) {
	//p.depth -= value
	p.aim -= value
}

var actions = map[string]string{
	"forward": "Forward",
	"up":      "Up",
	"down":    "Down",
}

func ParseAction(action string) (string, int64) {
	s_value := strings.Split(action, " ")
	a := s_value[0]
	v, _ := strconv.ParseInt(s_value[1], 10, 0)
	return a, v
}

func CallAction(position *position, action string, value int64) {
	f := reflect.ValueOf(position).MethodByName(actions[action])
	f_params := make([]reflect.Value, 1)
	f_params[0] = reflect.ValueOf(value)
	f.Call(f_params)
}

func main() {
	//input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	input, err := input.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := position{hpos: 0, depth: 0, aim: 0}

	for _, line := range input {
		action, value := ParseAction(line)
		CallAction(&p, action, value)
	}
	fmt.Printf("H=%d,D=%d ==> %d\n", p.hpos, p.depth, p.hpos*p.depth)
}
