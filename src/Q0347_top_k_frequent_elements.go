package src

import (
	"math/rand"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	occurrences := make(map[int]int, len(nums))
	for _, num := range nums {
		value, _ := occurrences[num]
		occurrences[num] = value + 1
	}

	lists := make([][]int, 0, len(occurrences))
	for num, value := range occurrences {
		lists = append(lists, []int{num, value})
	}

	lists = qsort(lists, 0, len(lists) - 1, k)
	res := make([]int, 0, k)
	for i := 1; i <= k; i++ {
		res = append(res, lists[len(lists) - i][0])
	}
	return res
}

func qsort(lists [][]int, start, end, k int) [][]int {
	if k >= end - start + 1 {
		return lists
	}
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int() % (end - start + 1) + start
	lists[start], lists[picked] = lists[picked], lists[start]

	index := start // 已筛选过的，频次小于等于当前lists[start]的数字
	for i := start + 1; i <= end; i++ {
		if lists[i][1] < lists[start][1] {
			lists[index + 1], lists[i] = lists[i], lists[index + 1]
			index++
		}
	}
	lists[index], lists[start] = lists[start], lists[index]

	if end - index + 1 < k {
		return qsort(lists, start, index - 1, k - (end - index + 1))
	} else if end - index + 1 > k {
		return qsort(lists, index + 1, end, k)
	}
	return lists
}
