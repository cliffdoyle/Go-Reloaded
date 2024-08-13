package utilities

import (
	"go-reloaded/utilities"
	"testing"
)

func TestToUpperComplete(t *testing.T) {
	result := utilities.UppercaseComplete("This is so so exciting (up, 2)")

	expected := "This is so SO EXCITING"

	if result != expected {
		t.Errorf("UppercaseComplete(\"This is so so exciting (up, 2)\") = %s; want %s", result, expected)
	}
}
