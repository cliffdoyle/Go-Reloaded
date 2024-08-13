package utilities

import (
	"testing"

	"go-reloaded/utilities"
)

func TestToLower(t *testing.T) {
	result := utilities.LowerCase("I should stop SHOUTING (low)")

	expected := "I should stop shouting"

	if result != expected {
		t.Errorf("LowerCase(\"I should stop SHOUTING (low)\") = %s; Found: %s", result, expected)
	}
}
