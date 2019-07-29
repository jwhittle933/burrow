package main

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)


// ReadLine ...
type ReadLine string

func (r ReadLine) toString() string {
	return string(r)
}

// Evaluate ...
func Evaluate(txt ReadLine) {
	if txt.toString() == "" {
		fmt.Println(aurora.Magenta("nil"))
	}
	parseExpression(txt)
}

func parseExpression(txt ReadLine) string {
	args := strings.Split(txt.toString(), " ")

	switch args[0]{
	case "func":
		// var nameAndArgs string
		// var returnType string
		// read func signature and compose func body
		// nameAndArgs = args[1]
		// returnType = args[2]
		return "func"
	case "if":
		// read condition and compose body
		// condition := args[1]
		return "if"
	case "for":
		// read loop conditions and execute
		return "for"
	case "var":
		// read assignment and save
		Assign(args[1], args[2:])
		return "assign"
	case "const":
		// read assignment and save
		return "assignC"
	default:
		return "assign"
	}

	return ""
}


func evalFunction(txt ReadLine) {
	//
}

func evalIf(txt ReadLine) {
	//
}

func evalInt(txt ReadLine) {
	//
}

func shouldExit(txt ReadLine) bool {
	if strings.Compare("exit", txt.toString()) == 0 {
		return true
	}
	return false
}

func replace(txt string) ReadLine {
	return ReadLine(strings.Replace(txt, "\n", "", -1))
}
