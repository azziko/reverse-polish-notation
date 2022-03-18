package main

import (
	"errors"
)

func parser(value string) (float64, error) {
	hierarchy := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	opsPool := Stack{}
	postfix := Queue{}
	currNumber := ""

	for key, v := range value {
		switch v {
		case '+', '-', '*', '/':
			//check if we have any unenqued number.
			//Enqueue in case we do
			//Or return an error because operator can't go first.
			if currNumber != "" {
				postfix.Enqueue(currNumber)
			} else {
				var err error
				err = errors.New("Invalid format")
				return 0, err
			}
			currNumber = "" //discard

			l := len(opsPool.ops)
			//check top stack value for hierarchy
			if l > 0 {
				for i := l - 1; i >= 0 && (hierarchy[opsPool.ops[i]] >= hierarchy[string(v)]); i-- {
					postfix.Enqueue(opsPool.Pop())
				}
				opsPool.Push(string(v))
			} else {
				opsPool.Push(string(v))
			}
		default:
			currNumber += string(v)
			//Enqueue the whole currNumber if we encounter the last key
			if key == len(value)-1 && currNumber != "" {
				postfix.Enqueue(currNumber)
			}
		}
	}
	//pop all operators remaining into the queue
	for len(opsPool.ops) > 0 {
		postfix.Enqueue(opsPool.Pop())
	}

	return parserRPN(postfix.RPN)
}
