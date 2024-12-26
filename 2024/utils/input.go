package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func ReadInputFile(target_filename string) string {
  // 0 caller - current input.go file, 1 caller - solution.go file
	_, filename, _, _ := runtime.Caller(1)
	scriptDir := filepath.Dir(filename)

	inputFilePath := filepath.Join(scriptDir, target_filename)
  return inputFilePath
}

func ReadLines(fileame string) ([]string, error) {
	file, err := os.Open(fileame)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	fmt.Println("Reading file: ", fileame)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ParseLines(lines []string) ([]int, []int) {
	var left []int
	var right []int
	for _, line := range lines {
		parts := strings.Fields(line)

		ints := make([]int, len(parts))

		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil, nil
			}
			ints[i] = num
		}
		left = append(left, ints[0])
		right = append(right, ints[1])
	}
	return left, right
}
