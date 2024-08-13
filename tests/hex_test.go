package utilities

import (
	"go-reloaded/utilities"
	"testing"
)

func TestHexadec(t *testing.T) {
	result := utilities.HexToDecimal("1E (hex) files were added")

	expected := "30 files were added"

	if result != expected {
		t.Errorf("HexToDecimal(\"1E (hex) files were added\") = %s; Found: %s", result, expected)
	}
}
