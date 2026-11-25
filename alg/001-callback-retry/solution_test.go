package callbackretry

import (
	"reflect"
	"testing"
)

func TestIncomplete(t *testing.T) {
	tests := []struct {
		name string
		logs []Log
		want []int
	}{
		{
			name: "문제 예시",
			logs: []Log{{3, "fail"}, {1, "fail"}, {3, "ok"}, {2, "fail"}, {1, "fail"}, {3, "fail"}, {2, "ok"}},
			want: []int{1},
		},
		{
			name: "로그 없음",
			logs: nil,
			want: []int{},
		},
		{
			name: "전부 완료",
			logs: []Log{{1, "ok"}, {2, "fail"}, {2, "ok"}},
			want: []int{},
		},
		{
			name: "등장 순서 유지",
			logs: []Log{{5, "fail"}, {2, "fail"}, {9, "fail"}, {2, "ok"}, {5, "fail"}},
			want: []int{5, 9},
		},
		{
			name: "ok 뒤 fail은 무시",
			logs: []Log{{7, "ok"}, {7, "fail"}, {7, "fail"}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Incomplete(tt.logs)
			if got == nil {
				got = []int{}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Incomplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIncomplete(b *testing.B) {
	logs := make([]Log, 1_000_000)
	for i := range logs {
		status := "fail"
		if i%3 == 0 {
			status = "ok"
		}
		logs[i] = Log{JobID: i % 200_000, Status: status}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Incomplete(logs)
	}
}
