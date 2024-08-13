package utilities

import (
	"go-reloaded/utilities"
	"testing"
)

func TestConvertAtoAn(t *testing.T) {
	result := utilities.AtoAn("There it was. A amazing rock!")

	expected := "There it was. An amazing rock!"

	if result != expected {
		t.Errorf("AtoAn(\"There it was. A amazing rock!\") = %s; Found: %s", result, expected)
	}
}
