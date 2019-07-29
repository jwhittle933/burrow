package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// Assign ...
func Assign(name string, args []string) {
	// args may contain "= someValue" or just a type definition
	if args[0] == "=" {
		variables[name] = args[1]
		fmt.Println(aurora.BrightBlue(args[1]))
		return
	}

	if isType := types[args[1]]; !isType {
		fmt.Println(aurora.Red(fmt.Sprintf("%s is not type", args[1])))
		return
	}

	// set variable to nil value
	// TODO: associate var, nil value, and type
	variables[name] = nil
}

// AssignC ...
func AssignC(name string, value interface{}) {
	//
}
