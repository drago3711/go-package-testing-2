package bar

import (
	"fmt"
	"github.com/drago3711/go-package-testing-1/pkg/foo"
)

func Print() {
	fmt.Println("bar")

}

func PrintFoo() {
	foo.Print()
}
