package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"os"
)

// ReadLine ...
type ReadLine string

// NewReadLine ...
func NewReadLine(txt string, err error) ReadLine {
	if err != nil {
		logger.E("Error reading input: %s", err.Error())
		os.Exit(1)
	}
	return ReadLine(strings.Replace(txt, "\n", "", -1)).trimWhite()
}

func (r ReadLine) toString() string {
	return string(r)
}

func (r ReadLine) trimWhite() ReadLine {
	s := strings.TrimSpace(r.toString())
	return ReadLine(s)
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
	case "package":
		logger.W("Packages are not supported in the REPL")
		return
	default:
		// check for numeric string and parse ensuing operations
		if num, err := strconv.ParseFloat(args[0], 64); err == nil {
			ParseArithmetic(num, args[1:])
			return
		}

		// increment and decrement operations do not have a space
		// after the initial integer, i.e., 10++
		// Check for "n++" or "n--" pattern and execute operation
		// if value is a inc. or dec. variable, i.e., x++,
		// isIncDecOp will return error and block will be skipped
		// TODO: allow for inc/dec of variables
		if num, op, err := isIncDecOp(args[0]); err == nil {
			num, _ := strconv.ParseFloat(num, 64)
			ParseArithmetic(num, op)
			return
		}

		if isReassign(args) {
			ReAssign(args[0], args[1:])
			return
		}

		// check for variable refernce operations, i.e. string concat, arithmetic, etc.

		// fmt.Println(aurora.Red(fmt.Sprintf("Unexpected %s", args[0])))
		logger.W("Unexpected %s", args[0])
		return
	}
}

func isIncDecOp(txt string) (string, []string, error) {
	// if variable name contains numeric, func will fail
	// also strips out decimals

	// Check for "n++" or "n--" pattern and execute operation
	regNum, _ := regexp.Compile(`[0-9].`)
	regIncDec, _ := regexp.Compile("[^0-9][+|-]")

	n := regNum.FindAllString(txt, -1)
	op := regIncDec.FindAllString(txt, -1)

	if len(n) == 0 || len(op) == 0 {
		return "", nil, errors.New("No match")
	}

	if len(n) > 1 || len(op) > 1 {
		return "", nil, fmt.Errorf("Found %v numbers and %v operators", len(n), len(op))
	}

	return n[0], op, nil
}

func isReassign(a []string) bool{
	if len(a) < 2 {
		return false
	}
	if a[1] == "=" {
		if variables[a[0]] != nil || constants[a[0]] != nil {
			return true
		}
		return false
	}
	return false
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
