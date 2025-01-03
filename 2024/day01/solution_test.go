package day01

import (
	"testing"

	"github.com/larrypozdeev/advent-of-code/2024/utils"
)

func TestPart1(t *testing.T) {
	data := []string{"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3"}
	parsed := utils.ParseColumns(data)

	expected := 11

	if result := Part1(parsed[0], parsed[1]); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	data := []string{"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3"}
	parsed := utils.ParseColumns(data)

	expected := 31

	if result := Part2(parsed[0], parsed[1]); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
