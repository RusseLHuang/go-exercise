package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestStringCalculatorAddEmptyString(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum := stringCalculator.Add("")

	if sum != 0 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddOneElement(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum := stringCalculator.Add("1")

	if sum != 1 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddSum(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum := stringCalculator.Add("1,2")

	if sum != 3 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddSupportNewLine(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum := stringCalculator.Add("1\n2,3\n5")

	if sum != 11 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddSupportDifferentDelimiters(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum := stringCalculator.Add("//;\n1;2;3")

	if sum != 6 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddInvalidNumber(t *testing.T) {
	stringCalculator := StringCalculator{}

	assert.Panics(t, func() { stringCalculator.Add("1,x") }, "Should panic when input is not number")
}
