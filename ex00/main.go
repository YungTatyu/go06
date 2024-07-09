package main

import (
	"ft"
	"os"
)

const (
	yes     = 1
	no      = 0
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

var lengthOfArg int = 0

func countArgc() int {
	var argv []string = os.Args[1:]
	var length int = 0
	for range argv {
		length++
	}
	return length
}

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) int {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func init() {
	lengthOfArg = countArgc()
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
