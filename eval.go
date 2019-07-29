package main

import (
	"fmt"
	"strings"
	"os"

	"github.com/logrusorgru/aurora"
)

// ReadLine ...
type ReadLine string

// NewReadLine ...
func NewReadLine(txt string, err error) ReadLine {
	if err != nil {
		fmt.Println(fmt.Sprintf("Error reading input; %s", aurora.Red(err.Error())))
		os.Exit(1)
	}
	return ReadLine(strings.Replace(txt, "\n", "", -1))
}

func (r ReadLine) toString() string {
	return string(r)
}

// Evaluate ...
func Evaluate(txt ReadLine) string {
	args := strings.Split(txt.toString(), " ")

	switch args[0] {
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
		AssignC(args[1], args[2:])
		return "assignC"
	case "import":
		return "import"
	case "pacakge":
		return "package"
	default:
		fmt.Println(aurora.Red(fmt.Sprintf("Unknown reference %s", args[0])))
		return ""
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
