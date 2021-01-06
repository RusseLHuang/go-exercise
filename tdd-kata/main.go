package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func (sc StringCalculator) Add(numStr string) (int, error) {
	sum := 0
	if numStr == "" {
		return sum, nil
	}

	numListString := numStr
	defaultDelimiter := ","
	if isCustomDelimiter(numStr) {
		defaultDelimiter = numStr[2:3]
		numListString = strings.Replace(numStr[3:], "\n", "", 1)
	}

	replacedStr := strings.ReplaceAll(numListString, "\n", ",")
	nums := strings.Split(replacedStr, defaultDelimiter)

	for i := 0; i < len(nums); i++ {
		res, _ := strconv.Atoi(nums[i])

		if res < 0 {
			return 0, fmt.Errorf("negatives not allowed")
		}

		sum += res
	}

	return sum, nil
}

func isCustomDelimiter(str string) bool {
	line := strings.Split(str, "\n")
	delimiterLine := line[0]

	return len(delimiterLine) >= 2 && delimiterLine[0:2] == "//"
}
