package main

import "fmt"

/**
45. 跳跃游戏 II

给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2


提示:

1 <= nums.length <= 10^4
0 <= nums[i] <= 1000
题目保证可以到达 nums[n-1]

*/

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{2, 3, 0, 1, 4}
	//nums := []int{1, 2, 1, 1, 1}
	result := jump2(nums)
	fmt.Println(result)
}

// 反向查找出发位置
func jump1(nums []int) int {
	position := len(nums) - 1
	step := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if nums[i]+i >= position {
				position = i
				step++
				break
			}
		}
	}
	return step
}

// 正向查找可到达的最大位置
func jump2(nums []int) int {
	maxIndex, count, next := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxIndex = max(maxIndex, i+nums[i])
		if i == next {
			next = maxIndex
			count++
		}
	}
	return count
}
