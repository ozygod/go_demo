package main

import "fmt"

/*
238. 除自身以外数组的乘积

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]


提示：

2 <= nums.length <= 10^5
-30 <= nums[i] <= 30
输入 保证 数组 answer[i] 在  32 位 整数范围内


进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
*/

func main() {
	//nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	results := productExceptSelf4(nums)
	fmt.Println(results)
}

// 暴力解法1
func productExceptSelf1(nums []int) []int {
	results := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result := 1
		for j := 0; j < len(nums); j++ {
			if j != i {
				result *= nums[j]
			}
		}
		results[i] = result
	}

	return results
}

// 暴力解法2
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	results := make([]int, n)
	left, right := 0, 0
	for i := 0; i < n; i++ {
		results[i] = 1
	}
	for left < n {
		if left != right {
			results[left] *= nums[right]
		}
		right++
		if right >= n {
			left++
			right = 0
		}
	}

	return results
}

// 方法一：左右乘积列表
func productExceptSelf3(nums []int) []int {
	n := len(nums)
	left, right, results := make([]int, n), make([]int, n), make([]int, n)

	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		results[i] = left[i] * right[i]
	}

	return results
}

// 空间复杂度 O(1) 的方法
func productExceptSelf4(nums []int) []int {
	n := len(nums)
	results := make([]int, n)

	results[0] = 1
	for i := 1; i < n; i++ {
		results[i] = results[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		results[i] = results[i] * right
		right *= nums[i]
	}

	return results
}
