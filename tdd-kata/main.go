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

	replacedStr := strings.ReplaceAll(numStr, "\n", ",")
	nums := strings.Split(replacedStr, ",")

	for i := 0; i < len(nums); i++ {
		res, err := strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}

		sum += res
	}

	return sum
}
