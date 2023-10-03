package day3

import "fmt"

func numIdenticalPairs(nums []int) int {
	pair := 0
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for _, v := range nums {
		val := m[v]
		pair += val - 1
		m[v]--
	}

	return pair
}

type TestS struct {
	nums []int
	ans  int
}

func Test3() {
	t := []TestS{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3, 4}, 0},
	}

	for i, v := range t {
		fmt.Printf("%d case %d : %d : %d \n", i, v.nums, v.ans, numIdenticalPairs(v.nums))
	}
}
