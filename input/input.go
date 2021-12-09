package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) ([]string, error) {
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

func ConvertStringToSliceOfInts(input string, separator string) ([]int, error) {
	ints := []int{}
	elements := strings.Split(input, separator)
	for _, el := range elements {
		n, err := strconv.ParseInt(el, 10, 0)
		if err != nil {
			return nil, err
		}
		ints = append(ints, int(n))
	}
	return ints, nil
}
