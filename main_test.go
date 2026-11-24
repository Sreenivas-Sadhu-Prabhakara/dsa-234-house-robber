package main

import "testing"

func TestRob(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{5}, 5},
		{[]int{}, 0},
		{[]int{2, 1}, 2},
		{[]int{1, 3, 1, 3, 100}, 103},
		{[]int{0, 0, 0}, 0},
	}
	for _, c := range cases {
		if got := Rob(c.in); got != c.want {
			t.Errorf("Rob(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
