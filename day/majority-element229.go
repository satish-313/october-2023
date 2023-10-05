package day

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) []int {
	res := []int{}

	val := float64(len(nums)) / 3.0
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		count := 1
		for j := i + 1; j < len(nums) && nums[i] == nums[j]; j++ {
			count += 1
			i += 1
		}
		if float64(count) > val {
			res = append(res, nums[i])
		}
	}
	return res
}

func Test5() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
