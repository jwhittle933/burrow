package main

import (
	"bufio"
	"fmt"
	"go/build"
	"io"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora"
)


func init() {
	printInfo()
	variables = make(map[string]interface{})
	constants = make(map[string]interface{})
	functions = make(map[string]func())
	types = make(map[string]bool)
}

func main() {
	reader := reader(os.Stdin)
	loops := 1
loop:
	for {
		fmt.Print(fmt.Sprintf("go(%v)> ", loops))
		txt, _ := reader.ReadString('\n')
		rTxt := replace(txt)
		if shouldExit(rTxt) {
			break loop
		}
		loops++
		Evaluate(rTxt)
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


