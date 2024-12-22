package calculator

import (
	"errors"
	"fmt"
	"strconv"
)

func isNum(elem string) bool {
	if _, err := strconv.ParseFloat(elem, 64); err == nil {
		return true
	}
	return false
}

func isOp(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func Calc(expression string) (float64, error) {
	elements := parse(expression)
	postfix, err := rewriteToPostfix(elements)
	if err != nil {
		return 0, err
	}
	return calculatePostfix(postfix)
}

func parse(str string) []string {
	var elems []string
	var cur string
	for _, char := range str {
		switch char {
		case '+', '-', '*', '/', '(', ')':
			if len(cur) > 0 {
				elems = append(elems, cur)
				cur = ""
			}
			elems = append(elems, string(char))
		case ' ':
			continue
		default:
			cur = cur + string(char)
		}
	}
	if len(cur) > 0 {
		elems = append(elems, cur)
	}
	return elems
}

func rewriteToPostfix(elems []string) ([]string, error) {
	var result []string
	var ops []string
	for _, elem := range elems {
		if isNum(elem) {
			result = append(result, elem)
		} else if isOp(elem) {
			for len(ops) > 0 && priority(ops[len(ops)-1], elem) {
				result = append(result, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, elem)
		} else if elem == "(" {
			ops = append(ops, elem)
		} else if elem == ")" {
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				result = append(result, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			if len(ops) == 0 {
				return nil, errors.New("no matching bracket")
			}
			ops = ops[:len(ops)-1]
		} else {
			return nil, fmt.Errorf("invalid expression")
		}
	}
	for len(ops) > 0 {
		if ops[len(ops)-1] == "(" {
			return nil, errors.New("no matching bracket")
		}
		result = append(result, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	return result, nil
}

func calculatePostfix(postfix []string) (float64, error) {
	var stack []float64
	for _, elem := range postfix {
		if isNum(elem) {
			num, _ := strconv.ParseFloat(elem, 64)
			stack = append(stack, num)
		} else if isOp(elem) {
			if len(stack) <= 1 {
				return 0, errors.New("invalid expression")
			}
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			switch elem {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("unknown operator: %s", elem)
			}
		} else {
			return 0, fmt.Errorf("invalid elem: %s", elem)
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}

func priority(op1, op2 string) bool {
	var prior1 int
	var prior2 int
	switch op1 {
	case "*", "/":
		prior1 = 2
	case "+", "-":
		prior1 = 1
	default:
		prior1 = 0
	}
	switch op2 {
	case "*", "/":
		prior2 = 2
	case "+", "-":
		prior2 = 1
	default:
		prior2 = 0
	}
	return prior1 >= prior2
}
