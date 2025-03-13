package main

import "fmt"

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
//
//假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
//
//更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
//返回 k。

func main() {
	var nums = []int{3, 2, 2, 3}
	k := removeElement3(nums, 3)
	fmt.Println(nums, k)

	var nums2 = []int{1, 2, 3, 4, 5}
	k2 := removeElement3(nums2, 1)
	fmt.Println(nums2, k2)
}

func removeElement1(nums []int, val int) int {
	result, p1, p2 := 0, 0, len(nums)-1
	for p1 <= p2 {
		if nums[p1] != val {
			p1++
			result++
			continue
		}
		if nums[p2] != val {
			nums[p1] = nums[p2]
			p1++
			result++
		}

		//nums[p2] = 0
		p2--
	}
	return result
}

func removeElement2(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func removeElement3(nums []int, val int) int {
	p1, p2 := 0, len(nums)-1
	for p1 <= p2 {
		if nums[p1] == val {
			nums[p1] = nums[p2]
			p2--
		} else {
			p1++
		}
	}
	return p1
}
