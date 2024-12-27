package utils

import (
	"testing"
)

func TestAbs(t *testing.T) {
	if Abs(-1) != 1 {
		t.Error("Abs(-1) should be 1")
	}
}
