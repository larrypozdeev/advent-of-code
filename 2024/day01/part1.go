package day01

import (
	"fmt"
	"path/filepath"
	"runtime"
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
	// Get the directory of the source file at runtime
	_, filename, _, _ := runtime.Caller(0) // The current file's full path
	scriptDir := filepath.Dir(filename)

	// Construct the path to input.txt relative to the script's location
	inputFilePath := filepath.Join(scriptDir, "input.txt")

	// Read and process the input file
	inputLines, err := utils.ReadLines(inputFilePath)
	if err != nil {
		panic(err)
	}
	left, right := utils.ParseLines(inputLines)
	part1_result := Part1(left, right)

	fmt.Println("Part 1 result:", part1_result)
}
