package utilities

import (
	"go-reloaded/utilities"
	"testing"
)

func TestCapitalize(t *testing.T) {
	result := utilities.CapitalizePrevious("Welcome to the Brooklyn bridge (cap)")

	expected := "Welcome to the Brooklyn Bridge"

	if result != expected {
		t.Errorf("CapitalizePrevious(\"1E (hex) files were added\") = %s; Found: %s", result, expected)
	}
}
