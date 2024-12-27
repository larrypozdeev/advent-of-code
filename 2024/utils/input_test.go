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
	lines := []string{"1  2", "3   4", "5 6"}
	left, right := ParseLines(lines)

	expectedLeft := []int{1, 3, 5}
	expectedRight := []int{2, 4, 6}

	if !reflect.DeepEqual(left, expectedLeft) {
		t.Errorf("Expected left %v, got %v", expectedLeft, left)
	}

	if !reflect.DeepEqual(right, expectedRight) {
		t.Errorf("Expected right %v, got %v", expectedRight, right)
	}
}

