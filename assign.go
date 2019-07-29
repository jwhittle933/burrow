package main

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Assign ...
func Assign(name string, args []string) {
	// args may contain "= someValue" or just a type definition
	if args[0] == "=" {
		variables[name] = map[string]string{
			"value": args[1],
			"type":  "string",
		}
		fmt.Println(aurora.BrightBlue(variables[name].(map[string]string)["value"]))
		return
	}

	if !isType(args[0]) {
		fmt.Println(aurora.Red(fmt.Sprintf("%s is not type", args[0])))
		return
	}

	// set variable to nil value
	// TODO: associate var, nil value, and type
	variables[name] = nil
}

func isType(t string) bool {
	for _, n := range types {
		if strings.Compare(t, n) == 0 {
			return true
		}
	}
	return false
}

// AssignC ...
func AssignC(name string, args []string) {
	if isConst := constants[name]; isConst != nil {
		fmt.Println(aurora.Red(fmt.Sprintf("%s is already a constant", name)))
		return
	}

	constants[name] = map[string]string{
		"value": args[2],
		"type":  args[0],
	}
}
