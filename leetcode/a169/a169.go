package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/**
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2


提示：
n == nums.length
1 <= n <= 5 * 10^4
-10^9 <= nums[i] <= 10^9
*/

func main() {
	//nums := []int{1, 1, 2}
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	n := majorityElement5(nums)
	fmt.Println(n)
}

// 哈希表
func majorityElement1(nums []int) int {
	maxNum := len(nums) / 2
	items := make(map[int]int)
	for _, num := range nums {
		items[num] += 1
		if items[num] > maxNum {
			return num
		}
	}

	return 0
}

// 排序
func majorityElement2(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}

// 随机抽取
func majorityElement3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxNum := len(nums) / 2

	fn := func(nums []int, target int) int {
		count := 0
		for i := 0; i < len(nums); i++ {
			if target == nums[i] {
				count++
			}
		}
		return count
	}

	for {
		current := nums[rand.Intn(len(nums)-1)]
		if fn(nums, current) > maxNum {
			return current
		}
	}
}

// 分治
func majorityElement4(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}

func majorityElementRec(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}

	mid := lo + (hi-lo)/2
	left := majorityElementRec(nums, lo, mid)
	right := majorityElementRec(nums, mid+1, hi)

	if left == right {
		return left
	}

	leftCount := countInRange(nums, left, lo, hi)
	rightCount := countInRange(nums, right, lo, hi)
	if leftCount > rightCount {
		return left
	} else {
		return right
	}
}

func countInRange(nums []int, target, lo, hi int) int {
	count := 0
	for i := lo; i < hi; i++ {
		if target == nums[i] {
			count++
		}
	}
	return count
}

// Boyer-Moore 投票算法
func majorityElement5(nums []int) int {
	current, count := nums[0], 0
	for _, num := range nums {
		if count == 0 {
			current = num
		}
		if current == num {
			count++
		} else {
			count--
		}
	}
	return current
}
