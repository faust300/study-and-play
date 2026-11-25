package sortstruct

import (
	"reflect"
	"testing"
)

func TestRankMentors(t *testing.T) {
	in := []Mentor{{"b", 100}, {"a", 300}, {"c", 100}, {"d", 500}}
	orig := append([]Mentor(nil), in...)
	got := RankMentors(in)
	want := []Mentor{{"d", 500}, {"a", 300}, {"b", 100}, {"c", 100}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RankMentors = %v, want %v", got, want)
	}
	if !reflect.DeepEqual(in, orig) {
		t.Errorf("입력이 바뀌었다: %v (원래 %v)", in, orig)
	}
	if got := RankMentors(nil); len(got) != 0 {
		t.Errorf("RankMentors(nil) = %v, want empty", got)
	}
}

func TestTopK(t *testing.T) {
	cases := []struct {
		words []string
		k     int
		want  []string
	}{
		{[]string{"go", "ts", "go", "rust", "ts", "go"}, 2, []string{"go", "ts"}},
		{[]string{"b", "a", "c", "a", "b", "c"}, 2, []string{"a", "b"}},
		{[]string{"x"}, 3, []string{"x"}},
		{nil, 2, []string{}},
	}
	for _, c := range cases {
		got := TopK(c.words, c.k)
		if got == nil {
			got = []string{}
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("TopK(%v, %d) = %v, want %v", c.words, c.k, got, c.want)
		}
	}
}

func TestMergeWindows(t *testing.T) {
	cases := []struct{ in, want []Window }{
		{[]Window{{8, 10}, {1, 3}, {2, 6}, {6, 7}}, []Window{{1, 7}, {8, 10}}},
		{[]Window{{1, 3}, {3, 5}}, []Window{{1, 5}}},
		{[]Window{{1, 10}, {2, 3}, {4, 5}}, []Window{{1, 10}}},
		{[]Window{{5, 6}, {1, 2}}, []Window{{1, 2}, {5, 6}}},
		{nil, []Window{}},
	}
	for _, c := range cases {
		got := MergeWindows(c.in)
		if got == nil {
			got = []Window{}
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeWindows(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
