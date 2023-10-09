package day

import "fmt"

func searchRange(nums []int, target int) []int {
	res := [2]int{-1, -1}

	if len(nums) == 0 {
		return res[:]
	}
	l, r := 0, len(nums)-1

	if nums[l] == target && nums[r] == target {
		res[0], res[1] = l, r
		return res[:]
	}

	if nums[l] == target {
		res[0] = 0
		res[1] = rightEnd(nums, target)
		return res[:]
	}

	if nums[r] == target {
		res[1] = r
		res[0] = leftEnd(nums, target)
		return res[:]
	}

	// left
	res[0] = leftEnd(nums, target)
	if res[0] == -1 {
		return res[:]
	}

	// right
	res[1] = rightEnd(nums, target)

	return res[:]
}

func rightEnd(nums []int, target int) int {
	l, m, r := 0, 0, len(nums)-1

	for l <= r {
		m = (l + r) / 2

		if nums[m] == target && nums[m+1] != target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func leftEnd(nums []int, target int) int {
	l, m, r := 0, 0, len(nums)-1

	for l <= r {
		m = (l + r) / 2

		if nums[m] == target && nums[m-1] != target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

type test9 struct {
	q []int
	t int
	a []int
}

func Test9() {
	t := []test9{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	}
	for _, v := range t {
		fmt.Println(searchRange(v.q, v.t), " = ", v.a)
	}
}
