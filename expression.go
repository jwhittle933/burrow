package main

// Expression ...
type Expression interface {
	funcType
}

type funcType interface {
	Caller(int)
}



