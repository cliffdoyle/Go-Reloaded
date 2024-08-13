package utilities

import (
	"testing"

	"go-reloaded/utilities"
)

func TestFormattext(t *testing.T) {
	result := utilities.FormatText("I was sitting over there ,and then BAMM !!")

	expected := "I was sitting over there, and then BAMM!!"

	if result != expected {
		t.Errorf("FormatText(\"I was sitting over there ,and then BAMM !!\") = %s; Found: %s", result, expected)
	}
}

func TestSingleQuote(t *testing.T) {
	result := utilities.Quotation("As Elton John said: ' I am the most well-known homosexual in the world '")

	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"

	if result != expected {
		t.Errorf("Quotation(\"As Elton John said: ' I am the most well-known homosexual in the world '\") = %s; Found: %s", result, expected)
	}
}