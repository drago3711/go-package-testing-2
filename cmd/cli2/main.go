package main

import (
	"fmt"
	"github.com/drago3711/go-package-testing-2/pkg/bar"
)

func main() {
	fmt.Println("Starting cmd2")
	bar.Print()
	bar.PrintFoo()
}
