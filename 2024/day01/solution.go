package day01

import (
	"fmt"
	"sort"

	"github.com/larrypozdeev/advent-of-code/2024/utils"
)

func Part1(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	diff := 0
	for i := 0; i < len(left); i++ {
		diff += utils.Abs(left[i] - right[i])
	}
	return diff
}

func Part2(left []int, right []int) int {
	leftOccurences := make(map[int]int)
	rightOccurences := make(map[int]int)

	total := 0
	if len(left) != len(right) {
		panic("Left and right arrays have different lengths")
	}

	for i := 0; i < len(left); i++ {
		leftOccurences[left[i]] += 1
		rightOccurences[right[i]] += 1
	}

  for key, leftCount := range leftOccurences {
    total += key * leftCount * rightOccurences[key]
  }

	return total
}

func Run() {
	inputFilePath := utils.ReadInputFile("input.txt")

	inputLines, err := utils.ReadLines(inputFilePath)
	if err != nil {
		panic(err)
	}
	left, right := utils.ParseLines(inputLines)

	part1Result := Part1(left, right)
	part2Result := Part2(left, right)

	fmt.Println("Part 1 result:", part1Result)
	fmt.Println("Part 2 result:", part2Result)
}

