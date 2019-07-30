package main

import (
	"bufio"
	"fmt"
	"go/build"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/logrusorgru/aurora"
)

var (
	variables map[string]interface{}
	// constants:
	// c := map[string]map[string]string{
	//      "test": map[string]string{
	//            "value": "hello",
	//            "type": "string",
	//      }
	// }
	constants map[string]interface{}
	functions map[string]func()
	types     []string

	logger *Log
)

func init() {
	printInfo()
	variables = make(map[string]interface{})
	constants = make(map[string]interface{})
	functions = make(map[string]func())
	types = []string{"bool", "byte", "complex128", "complex64", "error", "float32", "float64", "int", "int16", "int32", "int64", "int8", "rune", "string", "uint", "uint16", "uint32", "uint64", "uint8", "uintptr"}
	logger = StartLog()
}

func main() {
	reader := reader(os.Stdin)
	loops := 1
app:
	for {
		fmt.Print(fmt.Sprintf("go(%v)> ", loops))
		rTxt := NewReadLine(reader.ReadString('\n'))

		if shouldExit(rTxt.toString()) {
			break app
		}

		if rTxt.toString() == "" {
			fmt.Println(aurora.Magenta("nil"))
			continue app
		}

		if storedVar := checkVar(rTxt.toString()); storedVar != nil {
			fmt.Println(aurora.BrightGreen(storedVar.(map[string]string)["value"]))
			continue app
		}

		if storedConst := checkConst(rTxt.toString()); storedConst != nil {
			fmt.Println(aurora.BrightGreen(storedConst.(map[string]string)["value"]))
			continue app
		}

		Evaluate(rTxt)
		loops++
	}
}

func printInfo() {
	fmt.Println(
		fmt.Sprintf("Go Version [%v] GOROOT [%v] GOPATH [%v]",
			aurora.BrightMagenta(runtime.Version()),
			aurora.BrightMagenta(runtime.GOROOT()),
			aurora.BrightMagenta(build.Default.GOPATH)))
	fmt.Println("Burrow Interactive Go REPL v 0.1 - type \"exit\" to quit\n")
}

func reader(r io.Reader) *bufio.Reader {
	return bufio.NewReader(r)
}

func shouldExit(txt string) bool {
	if strings.Compare("exit", txt) == 0 {
		return true
	}
	return false
}

func checkVar(txt string) interface{} {
	return variables[txt]
}

func checkConst(txt string) interface{} {
	return constants[txt]
}
