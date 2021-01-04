package main

import (
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func (sc StringCalculator) Add(numStr string) int {
	sum := 0
	if numStr == "" {
		return sum
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
		res, err := strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}

		sum += res
	}

	return sum
}

func isCustomDelimiter(str string) bool {
	line := strings.Split(str, "\n")
	delimiterLine := line[0]

	return len(delimiterLine) >= 2 && delimiterLine[0:2] == "//"
}
