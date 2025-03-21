package main

import (
	"fmt"
	"sort"
)

/*
274. H 指数

给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇
论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。


示例 1：

输入：citations = [3,0,6,1,5]
输出：3
解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。
示例 2：

输入：citations = [1,3,1]
输出：1


提示：

n == citations.length
1 <= n <= 5000
0 <= citations[i] <= 1000
*/

func main() {
	//citations := []int{3, 0, 6, 1, 5}
	citations := []int{1, 3, 1}
	//citations := []int{1, 2, 2}
	//citations := []int{100}
	h := hIndex3(citations)
	fmt.Println(h)
}

// 暴力解法
func hIndex1(citations []int) int {
	h, n := 0, len(citations)
	hFind := func(citations []int, i, n int) int {
		num, r := 0, citations[i]
		for j := 0; r > 0 && j < n; j++ {
			if citations[j] >= r {
				num++
			}
		}
		return min(num, r)
	}

	for i := 0; i < n; i++ {
		hTmp := hFind(citations, i, n)
		if hTmp > h {
			h = hTmp
		}
	}

	return h
}

// 排序
func hIndex2(citations []int) int {
	h, n := 0, len(citations)
	sort.Ints(citations)
	for i := n - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
		}
	}
	return h
}

// 计数排序
func hIndex3(citations []int) int {
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	tot := 0
	for i := n - 1; i >= 0; i-- {
		tot += counter[i]
		if tot >= i {
			return i
		}
	}

	return 0
}

// TODO 二分搜索，暂时无法理解
func hIndex4(citations []int) int {

	return 0
}
