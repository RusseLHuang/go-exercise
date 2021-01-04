package main

import "testing"

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
