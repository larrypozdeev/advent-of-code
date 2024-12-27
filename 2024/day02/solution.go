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

		prevNumber := 0
		isIncreasing := row[0] < row[1]
		failed := false

		for _, number := range row {
			if prevNumber == 0 {
				prevNumber = number
				continue
			}
			if number == prevNumber {
				failed = true
				break
			}
			if utils.Abs(number-prevNumber) > 3 {
				failed = true
				break
			}

			if isIncreasing && prevNumber > number {
				failed = true
				break
			} else if !isIncreasing && prevNumber < number {
				failed = true
				break
			}
			prevNumber = number
		}

		if !failed {
			safeCount += 1
		}
	}

	return safeCount, nil
}

func Part2([][]int) {}

func Run() {
	inputFilePath := utils.ReadInputFile("input.txt")

	inputLines, err := utils.ReadLines(inputFilePath)
	if err != nil {
		panic(err)
	}
	parsed := utils.ParseRows(inputLines)

	part1Result, _ := Part1(parsed)
	// part2Result := Part2(parsed[0], parsed[1])
	//
	fmt.Println("Part 1 result:", part1Result)
	// fmt.Println("Part 2 result:", part2Result)
}
