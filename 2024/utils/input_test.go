package utils

import (
	"os"
	"reflect"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Create a temporary file and write some lines to it
	tempFile := t.TempDir() + "/testfile.txt"
	content := "line1\nline2\nline3"
	err := os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	lines, err := ReadLines(tempFile)
	if err != nil {
		t.Errorf("Failed to read lines: %v", err)
	}

	expectedLines := []string{"line1", "line2", "line3"}
	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("Expected %v, got %v", expectedLines, lines)
	}
}

func TestParseLines(t *testing.T) {
  lines1 := []string{"1", "2", "3"}
  parsed1 := ParseColumns(lines1)
  expected1 := [][]int{{1, 2, 3}}
  if !reflect.DeepEqual(parsed1, expected1) {
    t.Errorf("Expected %v, got %v", expected1, parsed1)
  }

	lines2 := []string{"1  2", "3   4", "5 6"}
	parsed2 := ParseColumns(lines2)

	expected2 := [][]int{{1, 3, 5}, {2, 4, 6}}

	if !reflect.DeepEqual(parsed2, expected2) {
		t.Errorf("Expected %v, got %v", expected2, parsed2)
	}

	lines3 := []string{"1 2 3", "4 5 6", "7 8 9"}
  expected3 := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	parsed3 := ParseColumns(lines3)
	if !reflect.DeepEqual(parsed3, expected3) {
		t.Errorf("Expected %v, got %v", expected3, parsed3)
	}
}
