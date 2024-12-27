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

	if result, error := Part1(parsed); result != expected {
    if error != nil {
      t.Error(error)
    }
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	data := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	parsed := utils.ParseRows(data)

	expected := 4

	if result, error := Part2(parsed); result != expected {
    if error != nil {
      t.Error(error)
    }
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPart2More(t *testing.T) {
	data := []string{
		"2 4 6 8 10", // Safe (strictly increasing, diffs in [1..3])
		"10 7 4 1 0", // Safe (strictly decreasing, diffs in [-1..-3])
		"1 2 3 6 10", // Not safe initially (last diff = 4); safe if we remove "10"
		"10 9 7 4 0", // Not safe initially (last diff = -4); safe if we remove "0"
		"3 2 3 9",    // Remains unsafe no matter which single level is removed
		"5 5 5",      // Differences are 0; remains unsafe even after any single removal
		"1 4 7",      // Safe (all diffs = 3)
		"7 4 1",      // Safe (all diffs = -3)
	}

	parsed := utils.ParseRows(data)
	expected := 6

	if result, err := Part2(parsed); result != expected {
		if err != nil {
			t.Error(err)
		}
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
