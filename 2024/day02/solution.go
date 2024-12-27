package day02

import (
	"errors"
	"fmt"

	"github.com/larrypozdeev/advent-of-code/2024/utils"
)

func Part1(data [][]int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("No data")
	}
	if len(data[0]) == 0 {
		return 0, errors.New("No data")
	}

	safeCount := 0

	for _, row := range data {
		if len(row) < 2 {
			return 0, errors.New("Not enough data")
		}

		if isValid(row) {
			safeCount++
			continue
		}
	}

	return safeCount, nil
}

func Part2(data [][]int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("No data")
	}

	safeCount := 0

	for _, row := range data {
		if len(row) < 2 {
			return 0, errors.New("Not enough data")
		}

		if isValid(row) {
			safeCount++
			continue
		}

		// check without one level
		for i := 0; i < len(row); i++ {
			modified := append([]int{}, row[:i]...)
			modified = append(modified, row[i+1:]...)

			if isValid(modified) {
				safeCount++
				break
			}
		}
	}

	return safeCount, nil
}

func isValid(row []int) bool {
	if len(row) < 2 {
		return true
	}

	for i := 1; i < len(row); i++ {
		diff := utils.Abs(row[i] - row[i-1])
		if diff < 1 || diff > 3 || (row[i] == row[i-1]) {
			return false
		}

		// check for increasing/decreasing
		if i > 1 {
			if (row[i-1] < row[i]) != (row[i-2] < row[i-1]) {
				return false
			}
		}
	}

	return true
}

func Run() {
	inputFilePath := utils.ReadInputFile("input.txt")

	inputLines, err := utils.ReadLines(inputFilePath)
	if err != nil {
		panic(err)
	}
	parsed := utils.ParseRows(inputLines)

	part1Result, _ := Part1(parsed)
	part2Result, _ := Part2(parsed)
	//
	fmt.Println("Part 1 result:", part1Result)
	fmt.Println("Part 2 result:", part2Result)
}
