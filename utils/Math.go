package utils

import "strconv"

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func MCM(ints []int) int {
	if len(ints) == 0 {
		return 0
	}
	result := ints[0]
	for i := 1; i < len(ints); i++ {
		result = result * ints[i] / MCD(result, ints[i])
	}
	return result
}
