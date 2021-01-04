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

func TestStringCalculatorAddInvalidNumber(t *testing.T) {
	stringCalculator := StringCalculator{}

	assert.Panics(t, func() { stringCalculator.Add("1,x") }, "Should panic when input is not number")
}
