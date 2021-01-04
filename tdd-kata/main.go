package main

import (
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func (sc StringCalculator) Add(numStr string) int {
	sum := 0
	nums := strings.Split(numStr, ",")
	for i := 0; i < len(nums); i++ {
		res, _ := strconv.Atoi(nums[i])
		sum += res
	}

	return sum
}
