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

	nums := strings.Split(numStr, ",")
	for i := 0; i < len(nums); i++ {
		res, err := strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}

		sum += res
	}

	return sum
}
