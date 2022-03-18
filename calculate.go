package main

import (
	"errors"
	"strconv"
)

func popFromRPN(stack []float64) ([]float64, float64, float64) {
	var ab []float64
	stack, ab = stack[:len(stack)-2], stack[len(stack)-2:]
	return stack, ab[0], ab[1]
}

func parserRPN(input []string) (float64, error) {
	stack := []float64{}
	var a, b, value float64
	for _, i := range input {
		switch i {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, errors.New("Invalid postfix value")
			}
			stack, a, b = popFromRPN(stack)
			switch i {
			case "+":
				value = a + b
			case "-":
				value = a - b
			case "*":
				value = a * b
			case "/":
				value = a / b
			}
		default:
			var err error
			value, err = strconv.ParseFloat(i, 64)
			if err != nil {
				return 0, err
			}
		}
		stack = append(stack, value)
	}
	if len(stack) != 1 {
		return 0, errors.New("Inacceptable expression")
	}
	return stack[0], nil
}
