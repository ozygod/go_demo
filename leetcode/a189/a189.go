package main

import "fmt"

/**
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]


提示：

1 <= nums.length <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
0 <= k <= 10^5

进阶：

- 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
- 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
*/

func main() {
	//nums := []int{-1, -100, 3, 99}
	//k := 2
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3
	//nums := []int{1, 2}
	//k := 0
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 4
	rotate3(nums, k)
	fmt.Println(nums)
}

// 使用额外数组
func rotate1(nums []int, k int) {
	n := len(nums)
	if k > n {
		k = k % n
	}
	tmp, left := make([]int, n), n-k
	for i := 0; i < n; i++ {
		if i < left {
			tmp[i+k] = nums[i]
		} else {
			tmp[i-left] = nums[i]
		}
	}
	copy(nums, tmp)
}

// TODO 环状替换，未想明白
func rotate2(nums []int, k int) {
	n := len(nums)
	if k > n {
		k = k % n
	}
	if k == 0 {
		return
	}
	var x int
	for i := 0; i < k; i++ {
		x = (i + k) % n
		nums[i%k], nums[x] = nums[x], nums[i%k]
		fmt.Println(i, nums)
	}
}

// 翻转数组
func rotate3(nums []int, k int) {
	k = k % len(nums)
	reverse := func(nums []int) {
		for i, n := 0, len(nums); i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
