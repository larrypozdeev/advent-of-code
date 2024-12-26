package day01

import (
	"testing"

	"github.com/larrypozdeev/advent-of-code/2024/utils"
)

func TestExample(t *testing.T) {
	// Test input
	data := []string{"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3"}
	left, right := utils.ParseLines(data)

	expected := 11

	if result := Part1(left, right); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
