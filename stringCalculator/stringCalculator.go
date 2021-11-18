package stringCalculator

import (
	"strconv"
	"strings"
)

func SplitString(number string) (int, int) {
	numbersSplited := strings.Split(number, ",")
	a, err1 := strconv.Atoi(numbersSplited[0])
	b, err2 := strconv.Atoi(numbersSplited[1])
	if err1 == nil && err2 == nil {
		return a, b
	}
	return -1, -1
}
func Add(numbers string) int {

	a, b := SplitString(numbers)
	return a + b
}
