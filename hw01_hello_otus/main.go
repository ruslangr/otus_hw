package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	msg := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(msg))
}
