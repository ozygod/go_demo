package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 4
	ch1 <- 2
	ch1 <- 3

	for i := 0; i < cap(ch1); i++ {
		fmt.Println("element: ", <-ch1)
	}
}
