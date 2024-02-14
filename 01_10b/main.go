package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) (*float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid operand of type float64", op)
	}
	return &parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (*float64, error) {
	var allErrors error
	ops := strings.Fields(expr)
	if len(ops) != 3 {
		return nil, fmt.Errorf("cannot evaluate expression: need 3 ops, but got %d instead", len(ops))
	}
	left, e1 := parseOperand(ops[0])
	right, e2 := parseOperand(ops[2])

	opErr := errors.Join(e1, e2)
	if opErr != nil {
		allErrors = errors.Join(allErrors, opErr)
	}

	f := operators[ops[1]]
	if f == nil {
		allErrors = errors.Join(allErrors, fmt.Errorf("%s is an invalid operator, please use one of the following: + - * /", ops[1]))
	}

	if allErrors != nil {
		return nil, allErrors
	}

	result := f(*left, *right)
	return &result, nil
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()
	result, err := calculate(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s = %.2f\n", *expr, *result)
}
