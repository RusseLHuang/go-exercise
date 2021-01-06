package main

import (
	"testing"
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
	sum, _ := stringCalculator.Add("")

	if sum != 0 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddOneElement(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum, _ := stringCalculator.Add("1")

	if sum != 1 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddSum(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum, _ := stringCalculator.Add("1,2")

	if sum != 3 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddSupportNewLine(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum, _ := stringCalculator.Add("1\n2,3\n5")

	if sum != 11 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddSupportDifferentDelimiters(t *testing.T) {
	stringCalculator := StringCalculator{}
	sum, _ := stringCalculator.Add("//;\n1;2;3")

	if sum != 6 {
		t.Fatalf("Empty string should return zero")
	}
}

func TestStringCalculatorAddNegativeNumber(t *testing.T) {
	stringCalculator := StringCalculator{}
	_, err := stringCalculator.Add("1,-2")

	if err == nil || err.Error() != "negatives not allowed -2" {
		t.Fatalf("Should return error negatives not allowed")
	}
}

func TestStringCalculatorAddMultipleNegativeNumber(t *testing.T) {
	stringCalculator := StringCalculator{}
	_, err := stringCalculator.Add("1,-2,-3,-5")

	if err == nil || err.Error() != "negatives not allowed -2,-3,-5" {
		t.Fatalf("Should return error negatives not allowed")
	}
}

func TestStringCalculatorGetCalledCount(t *testing.T) {
	stringCalculator := StringCalculator{}
	_, err := stringCalculator.Add("1,-2,-3,-5")

	if err == nil || err.Error() != "negatives not allowed -2,-3,-5" {
		t.Fatalf("Should return error negatives not allowed")
	}

	_, err = stringCalculator.Add("1,-2")

	if err == nil || err.Error() != "negatives not allowed -2" {
		t.Fatalf("Should return error negatives not allowed")
	}

	count := stringCalculator.GetCalledCount()
	if count != 2 {
		t.Fatalf("Should called two times")
	}
}
