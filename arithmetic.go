package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/logrusorgru/aurora"
)

// ParseArithmetic ...
func ParseArithmetic(num float64, args []string) {

	if len(args) == 0 {
		fmt.Println(aurora.BrightCyan(fmt.Sprintf("%v", num)))
		return
	}


	arops := make(map[string]int)
	for _, item := range args {
		if item == "+" || item == "-" || item == "*" || item == "%" || item == "/" {
			_, exist := arops[item]
			if exist {
				arops[item]++
			} else {
				arops[item] = 1
			}
		}

	}

	switch args[0] {
	case "+":
		f, _ := strconv.ParseFloat(args[1], 64)
		fmt.Println(aurora.BrightCyan(add(num, f)))
		return
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

func add(one, two float64) float64 {
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
