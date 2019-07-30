package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
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
func Evaluate(txt ReadLine) {
	args := strings.Split(txt.toString(), " ")

	switch args[0] {
	case "func":
		// var nameAndArgs string
		// var returnType string
		// read func signature and compose func body
		// nameAndArgs = args[1]
		// returnType = args[2]
		return
	case "if":
		// read condition and compose body
		// condition := args[1]
		return
	case "for":
		// read loop conditions and execute
		return
	case "var":
		// read assignment and save
		Assign(args[1], args[2:])
	case "const":
		// read assignment and save
		AssignC(args[1], args[2:])
	case "import":
		return
	case "pacakge":
		return
	default:
		// check for arithmetic and perform operation
		if num, err := strconv.ParseFloat(getNumeric(args[0]), 64); err == nil {
			ParseArithmetic(num, args[1:])
			return
		}

		// check for variable refernce operations, i.e. string concat, arithmetic, etc.

		fmt.Println(aurora.Red(fmt.Sprintf("Unknown reference %s", args[0])))
		return
	}
}

func getNumeric(txt string) string {
	// strip off ++ or -- from string (or any other non-numeric value)
	// if variable name contains numeric, func will fail
	reg, _ := regexp.Compile("[^0-9]+")
	return reg.ReplaceAllString(txt, "")
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
