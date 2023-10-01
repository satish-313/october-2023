package day1

import (
	"strings"
)

func BestReverse(s string) string {
	words := strings.Split(s, " ")
	ret := []string{}

	for _, word := range words {
		temp := []byte(word)
		for i, j := 0, len(temp)-1; i < j; {
			temp[i], temp[j] = temp[j], temp[i]
			j--
			i++
		}
		ret = append(ret, string(temp))
	}

	return strings.Join(ret, " ")
}

func swap(m, n int, arr []string) {
	mid := (n - m) / 2

	for i := 0; i <= mid; i++ {
		temp := arr[m+i]
		arr[m+i] = arr[n-i]
		arr[n-i] = temp
	}
}

func Reverse(s string) string {
	l, len, arr := 0, len(s), strings.Split(s, "")

	for i, v := range s {
		if string(v) == " " {
			swap(l, i-1, arr)
			l = i + 1
		}
	}
	swap(l, len-1, arr)
	return strings.Join(arr, "")
}
