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

func RunPart1() {
  inputFilePath:=utils.ReadInputFile("input.txt")

	inputLines, err := utils.ReadLines(inputFilePath)
	if err != nil {
		panic(err)
	}
	left, right := utils.ParseLines(inputLines)
	part1_result := Part1(left, right)

	fmt.Println("Part 1 result:", part1_result)
}


func Part2(left []int, right []int) int {
	return 0
}

func RunPart2() {
}
