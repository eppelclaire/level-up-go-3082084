package main

import (
	"flag"
	"log"
	"regexp"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	regExp := regexp.MustCompile("[^(\\[\\(\\)\\]\\{\\})]*")
	newStr := regExp.ReplaceAllString(expr, "")
	//fmt.Println(newStr)

	if len(newStr)%2 != 0 {
		return false
	}

	pairings := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	var exprStack []string

	for _, c := range newStr {
		switch string(c) {
		case "(", "[", "{":
			exprStack = append(exprStack, string(c))
		case ")", "]", "}":
			lastChar := exprStack[len(exprStack)-1]
			if lastChar != pairings[string(c)] {
				return false
			}
			exprStack = exprStack[:len(exprStack)-1]
		}
	}

	return len(exprStack) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
