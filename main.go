package go_math

import "errors"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divided by zero operation is prohabited.")
	}
	return a / b, nil
}
