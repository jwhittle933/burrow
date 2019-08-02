package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

// Scan ...
func Scan(r ReadLine) {
	src := []byte(r)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Println(tok, lit)
	}
}
