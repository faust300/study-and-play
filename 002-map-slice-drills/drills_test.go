package drills

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	got := CountWords([]string{"a", "b", "a", "c", "a"})
	want := map[string]int{"a": 3, "b": 1, "c": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CountWords = %v, want %v", got, want)
	}
	if got := CountWords(nil); len(got) != 0 {
		t.Errorf("CountWords(nil) = %v, want empty", got)
	}
}

func TestDedupe(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{3, 1, 3, 2, 1}, []int{3, 1, 2}},
		{[]int{5, 5, 5}, []int{5}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		got := Dedupe(c.in)
		if got == nil {
			got = []int{}
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Dedupe(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestGroupByLength(t *testing.T) {
	got := GroupByLength([]string{"go", "ts", "rust", "js", "c"})
	want := map[int][]string{1: {"c"}, 2: {"go", "ts", "js"}, 4: {"rust"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupByLength = %v, want %v", got, want)
	}
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		wi, wj int
		wok    bool
	}{
		{[]int{2, 7, 11, 15}, 9, 0, 1, true},
		{[]int{3, 2, 4}, 6, 1, 2, true},
		{[]int{3, 3}, 6, 0, 1, true},
		{[]int{1, 2, 3}, 100, 0, 0, false},
	}
	for _, c := range cases {
		i, j, ok := TwoSum(c.nums, c.target)
		if ok != c.wok || (ok && (i != c.wi || j != c.wj)) {
			t.Errorf("TwoSum(%v, %d) = (%d, %d, %v), want (%d, %d, %v)", c.nums, c.target, i, j, ok, c.wi, c.wj, c.wok)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{7}, []int{7}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		Reverse(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("Reverse → %v, want %v", c.in, c.want)
		}
	}
}
