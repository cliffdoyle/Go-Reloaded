package utilities

import (
	"go-reloaded/utilities"
	"testing"
)

func TestBinary(t *testing.T) {
	result := utilities.BinarytoDecimal("It has been 10 (bin) years" )

	expected := "It has been 2 years"

	if result != expected {
		t.Errorf("BinarytoDecial(\"It has been 10 (bin) years\") = %s; Found: %s", result, expected)
	}
}
