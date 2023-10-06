package day

import (
	"fmt"
	"math"
)

func rec(i int, dp []int) int {
	if i <= 3 {
		return i
	}
	if dp[i] > 0 {
		return dp[i]
	}
	ans := i
	for j := 2; j < i; j++ {
		r := rec(i-j, dp) * rec(j, dp)
		ans = int(math.Max(float64(r), float64(ans)))
	}
	dp[i] = ans
	return ans
}

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	dp := [60]int{}

	return rec(n, dp[:])
}

type test struct {
	q int
	a int
}

func Test6() {
	t := []test{
		{2, 1},
		{10, 36},
	}

	for i, v := range t {
		fmt.Printf("%d case %d = %d\n", i, v.a, integerBreak(v.q))
	}
}
