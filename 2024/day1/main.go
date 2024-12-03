package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <input-file>")
		return -1
	}
	fileName := args[0]
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return -1
	}
	p1 := partOne(string(contents))
	fmt.Println("Part One:", p1)
	p2 := partTwo(string(contents))
	fmt.Println("Part Two:", p2)
	return 0
}

func prepList(input string) ([]int, []int) {
	input = strings.Replace(input, "   ", ",", -1)
	left, right := []int{}, []int{}
loop:
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			break loop
		}
		l, err := strconv.Atoi(parts[0])
		if err != nil {
			return []int{}, []int{}
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			return []int{}, []int{}
		}
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func partOne(input string) int {
	left, right := prepList(input)

	slices.Sort(left)
	slices.Sort(right)

	diffs := []int{}

	if len(left) != len(right) {
		return -1
	}

	for i := 0; i < len(left); i++ {
		d := left[i] - right[i]
		if d < 0 {
			d *= -1
		}
		diffs = append(diffs, d)
	}

	diff := sum(diffs)

	return diff
}

func partTwo(input string) int {
	left, right := prepList(input)
	rightMap := listToMap(right)

	distances := []int{}
	for _, l := range left {
		dist := l * rightMap[l]
		distances = append(distances, dist)
	}
	return sum(distances)
}

func listToMap(input []int) map[int]int {
	out := map[int]int{}
	for _, v := range input {
		out[v]++
	}
	return out
}

type Numeric interface {
	~int | ~float64 | ~float32
}

func sum[T Numeric](input []T) T {
	var sum T
	for _, v := range input {
		sum += v
	}
	return sum
}
