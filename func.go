package main

import (
	"strings"
)

// FuncBuilder ...
type FuncBuilder struct {
	Args []string
	Return []string
}

func isFunc(txt ReadLine) bool {
	args := strings.Split(txt.toString(), " ")
	if args[0] == "func" {
		return true
	}
	return false
}
