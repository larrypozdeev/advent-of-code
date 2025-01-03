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

func ParseColumns(lines []string) [][]int {
	if len(lines) == 0 {
		return nil
	}

	// parse the first line to determine the number of columns
	firstLineParts := strings.Fields(lines[0])
	numCols := len(firstLineParts)

	result := make([][]int, numCols)
	for i := range result {
		result[i] = make([]int, len(lines))
	}

	for rowIndex, line := range lines {
		parts := strings.Fields(line)
		for colIndex, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
			result[colIndex][rowIndex] = num
		}
	}

	return result
}

func ParseRows(lines []string) [][]int {
	if len(lines) == 0 {
		return nil
	}

	result := make([][]int, len(lines))

	for i := 0; i < len(lines); i++ {
		parts := strings.Fields(lines[i])
		result[i] = make([]int, len(parts))
		for j := 0; j < len(parts); j++ {
			num, err := strconv.Atoi(parts[j])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
			result[i][j] = num
		}
	}
	return result
}
