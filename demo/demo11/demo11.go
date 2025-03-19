package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Enter function main.")
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("panic: %s\n", err)
	//	}
	//}()
	//// 引发panic。
	//panic(errors.New("something wrong"))
	//fmt.Println("Exit function main.")
	something()

	// 在 defer 函数中引发 panic
	defer func() {
		if r := recover(); r != nil {
			// 打印出 panic 的信息
			fmt.Println("Recovered in defer:", r)
			// 在 defer 函数中引发新的 panic
			panic("New panic in defer")
		}
	}()

	// 引发 panic
	panic("Oops! Main panic!")
}

func something() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}
