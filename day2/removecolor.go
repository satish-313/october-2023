package day2

import "fmt"

func winnerOfGame(colors string) bool {
	ac := 0
	bc := 0

	for i := 1; i < len(colors)-1; i++ {
		p := colors[i-1]
		c := colors[i]
		n := colors[i+1]

		if c == p && c == n {
			if string(c) == "A" {
				ac += 1
			} else {
				bc += 1
			}
		}
	}

	if ac == 0 || ac <= bc {
		return false
	}
	return true
}

type testS struct {
	val string
	ans bool
}

func Testday2() {
	test := [3]testS{
		{"AAABABB", true},
		{"AA", false},
		{"ABBBBBBBAAA", false},
	}

	for i, v := range test {
		fmt.Printf("%d case %s ans is %t == %t\n", i, v.val, v.ans, winnerOfGame(v.val))
	}
}
