package main

import (
	"fmt"
	"math"

	"github.com/logrusorgru/aurora"
)

// ParseArithmetic ...
func ParseArithmetic(num float64, args []string) {

	if len(args) == 0 {
		fmt.Println(aurora.BrightCyan(fmt.Sprintf("%v: %T", num, num)))
		return
	}

	switch args[0] {
	case "+":
		fmt.Println("Addition")
	case "-":
		//
	case"*":
		//
	case "/":
		//
	case "%":
		//
	case "++": // ++ will be in num
		increment(num)
	case "--": // -- will be in num
		//
	default:
		fmt.Println(aurora.BrightCyan(fmt.Sprintf("%v: %T", num, num)))
	}
}

func add(one, two int) int {
	return one + two
}

func sub(one, two int) int {
	return one - two
}

func div(one, two float64) float64 {
	return math.Remainder(one, two)
}

func increment(num float64) {
	num++
	fmt.Println(aurora.BrightGreen(num))
}
