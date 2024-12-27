package day02

import (
	"testing"

	"github.com/larrypozdeev/advent-of-code/2024/utils"
)

func TestPart1(t *testing.T) {
	data := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	parsed := utils.ParseRows(data)

	expected := 2

	if result := Part1(parsed); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
}
