package main

import "fmt"

func main() {
	demo1()
	demo2()
	//demo3()
	demo4()
	demo5()
}

func demo1() {
	nums := [6]int{1, 2, 3, 4, 5, 6}
	maxLen := len(nums) - 1
	for i, v := range nums {
		if i == maxLen {
			nums[0] += v
		} else {
			nums[i+1] += v
		}
	}
	fmt.Println(nums)
}
func demo2() {
	nums := []int{1, 2, 3, 4, 5, 6}
	maxLen := len(nums) - 1
	for i, v := range nums {
		if i == maxLen {
			nums[0] += v
		} else {
			nums[i+1] += v
		}
	}
	fmt.Println(nums)
}

//	func demo3() {
//		nums := []int8{0, 1, 2, 3, 4, 5, 6}
//		switch 1 + 3 {
//		case nums[0], nums[1]:
//			fmt.Println("1 or 2")
//		case nums[2], nums[3]:
//			fmt.Println("2 or 3")
//		case nums[4], nums[5], nums[6]:
//			fmt.Println("4 or 5 or 6")
//		}
//	}

func demo4() {
	nums := []int8{0, 1, 2, 3, 4, 5, 6}
	switch nums[4] {
	case 0, 1:
		fmt.Println("1 or 2")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}

func demo5() {
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or 6")
	}
}

func demo6() {
	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	//case uint8, uint16:
	//	fmt.Println("uint8 or uint16")
	//case byte:
	//	fmt.Printf("byte")
	default:
		fmt.Printf("unsupported type: %T", t)
	}
}
