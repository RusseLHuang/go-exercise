package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type StringCalculator struct {
	AddCalledCount int
}

func (sc StringCalculator) GetCalledCount() int {
	return sc.AddCalledCount
}

func (sc *StringCalculator) Add(numStr string) (int, error) {
	sc.AddCalledCount++

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

	var negativesBuffer bytes.Buffer

	for i := 0; i < len(nums); i++ {
		res, _ := strconv.Atoi(nums[i])

		if res < 0 {
			negativesBuffer.WriteString(nums[i])
			negativesBuffer.WriteString(",")
		}

		sum += res
	}

	if negativesBuffer.String() != "" {
		negativesNums := negativesBuffer.String()
		return 0, fmt.Errorf("negatives not allowed %s", negativesNums[:len(negativesNums)-1])
	}

	return sum, nil
}

func isCustomDelimiter(str string) bool {
	line := strings.Split(str, "\n")
	delimiterLine := line[0]

	return len(delimiterLine) >= 2 && delimiterLine[0:2] == "//"
}
