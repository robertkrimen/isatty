package main

import (
	"fmt"
	"github.com/robertkrimen/isatty"
	"os"
)

func main() {
	result := isatty.Check(os.Stdin.Fd())
	fmt.Println(result)
	if !result {
		os.Exit(1)
	}
}
