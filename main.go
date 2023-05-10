package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("emacs")
}

func Foo() string {
	return os.Getenv("FOO")
}
