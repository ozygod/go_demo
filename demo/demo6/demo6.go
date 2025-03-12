package main

import "fmt"

type operator func(x, y int) int

func cal(x, y int, op operator) int {
	return op(x, y)
}

func main() {
	fmt.Println("+:", cal(1, 2, func(x, y int) int {
		return x + y
	}))
	fmt.Println("-:", cal(1, 2, func(x, y int) int {
		return x - y
	}))
	fmt.Println("*:", cal(1, 2, func(x, y int) int {
		return x * y
	}))
	fmt.Println("/:", cal(1, 2, func(x, y int) int {
		return x / y
	}))

	complexArray1 := [3][]string{[]string{"d", "e", "f"}, []string{"g", "h", "i"}, []string{"j", "k", "l"}}
	fmt.Println(complexArray1)
	copyDemo(complexArray1)
	fmt.Println(complexArray1)
}

func copyDemo(arr [3][]string) {
	arr[0] = []string{"a", "b", "c"}
	arr[1][0] = "gg"
}
