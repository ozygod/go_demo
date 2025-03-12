package main

import "fmt"

type Printer func(content string)

func print(content string) {
	fmt.Println(content)
}

func main() {
	var p Printer
	p = print
	p("hello world")
}
