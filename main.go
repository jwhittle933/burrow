package main

import (
	"bufio"
	"fmt"
	"go/build"
	"os"
	"runtime"
	"strings"

	"github.com/logrusorgru/aurora"
)

func main() {
	info()

	reader := reader()
	loops := 0
loop:
	for {
		fmt.Print(fmt.Sprintf("go(%v)> ", loops))
		txt, _ := reader.ReadString('\n')
		if shouldExit(replace(txt)) {
			break loop
		}
		loops ++
		if replace(txt) == "" {
			fmt.Println(aurora.Magenta("nil"))
		}
	}
}

func info() {
	fmt.Println("Burrow Interactive Go REPL v 0.1\n")
	fmt.Println(
		fmt.Sprintf("Version [%s] Root [%s] Path [%s]",
			runtime.Version(),
			runtime.GOROOT(),
			build.Default.GOPATH))
}

func reader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func shouldExit(txt string) bool {
	if strings.Compare("exit", txt) == 0 {
		return true
	}
	return false
}

func replace(txt string) string {
	return strings.Replace(txt, "\n", "", -1)
}
