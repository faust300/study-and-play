package parse

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	cases := []struct {
		line    string
		want    Log
		wantErr error
	}{
		{"3,ok", Log{3, "ok"}, nil},
		{" 7 , fail ", Log{7, "fail"}, nil},
		{"0,fail", Log{0, "fail"}, nil},
		{"3", Log{}, ErrFormat},
		{"1,ok,extra", Log{}, ErrFormat},
		{"", Log{}, ErrFormat},
		{"abc,ok", Log{}, ErrJobID},
		{"3.5,ok", Log{}, ErrJobID},
		{"3,done", Log{}, ErrStatus},
		{"3,OK", Log{}, ErrStatus},
	}
	for _, c := range cases {
		got, err := ParseLine(c.line)
		if c.wantErr != nil {
			if !errors.Is(err, c.wantErr) {
				t.Errorf("ParseLine(%q) err = %v, want %v", c.line, err, c.wantErr)
			}
			continue
		}
		if err != nil {
			t.Errorf("ParseLine(%q) unexpected err = %v", c.line, err)
			continue
		}
		if got != c.want {
			t.Errorf("ParseLine(%q) = %+v, want %+v", c.line, got, c.want)
		}
	}
}

func TestQueue(t *testing.T) {
	var q Queue
	if v, ok := q.Pop(); ok || v != 0 {
		t.Fatalf("빈 큐 Pop = (%d, %v), want (0, false)", v, ok)
	}
	q.Push(3)
	q.Push(1)
	q.Push(4)
	if q.Len() != 3 {
		t.Fatalf("Len = %d, want 3", q.Len())
	}
	for _, want := range []int{3, 1} {
		if v, ok := q.Pop(); !ok || v != want {
			t.Fatalf("Pop = (%d, %v), want (%d, true)", v, ok, want)
		}
	}
	q.Push(9)
	for _, want := range []int{4, 9} {
		if v, ok := q.Pop(); !ok || v != want {
			t.Fatalf("Pop = (%d, %v), want (%d, true)", v, ok, want)
		}
	}
	if q.Len() != 0 {
		t.Fatalf("Len = %d, want 0", q.Len())
	}
	if _, ok := q.Pop(); ok {
		t.Fatal("다 꺼낸 뒤 Pop의 ok가 true")
	}
}

func TestSummarize(t *testing.T) {
	cases := []struct {
		lines          []string
		wantIncomplete []int
		wantBad        []int
	}{
		{[]string{"3,fail", "x,ok", "1,fail", "3,ok", "2"}, []int{1}, []int{1, 4}},
		{[]string{"1,ok", "2,fail"}, []int{2}, []int{}},
		{[]string{"bad", "", "1,done"}, []int{}, []int{0, 1, 2}},
		{nil, []int{}, []int{}},
	}
	for _, c := range cases {
		inc, bad := Summarize(c.lines)
		if inc == nil {
			inc = []int{}
		}
		if bad == nil {
			bad = []int{}
		}
		if !reflect.DeepEqual(inc, c.wantIncomplete) || !reflect.DeepEqual(bad, c.wantBad) {
			t.Errorf("Summarize(%q) = (%v, %v), want (%v, %v)", c.lines, inc, bad, c.wantIncomplete, c.wantBad)
		}
	}
}
