// Package callbackretry — 001. 콜백 재시도 로그
//
// 문제
//
// ML 서버가 보낸 콜백 로그가 시간순으로 쌓여 있다. 로그 한 줄은 (job_id, status)이고
// status는 "fail" 또는 "ok"다. 같은 job_id는 여러 번 재시도될 수 있다.
// 어떤 작업이든 "ok"가 한 번이라도 나오면 완료로 보고, 그 뒤 같은 작업의 로그는 무시한다.
// 완료되지 못한 작업의 job_id를 처음 등장한 순서대로 반환하라.
//
// 예시
//
//	입력: [(3,fail) (1,fail) (3,ok) (2,fail) (1,fail) (3,fail) (2,ok)]
//	출력: [1]
//
//	3은 세 번째 로그에서 ok → 완료 (뒤의 fail은 무시)
//	1은 fail만 두 번 → 미완료
//	2는 마지막 로그에서 ok → 완료
//
// 제약
//
//	로그 수 N ≤ 10^6, O(N) 풀이를 목표로 한다.
//
// 연습하는 Go 문법: :=, for·range, _, map[int]bool, append
package callbackretry

// Log는 콜백 로그 한 줄이다. Status는 "fail" 또는 "ok".
type Log struct {
	JobID  int
	Status string
}

// Incomplete는 완료되지 못한 작업의 job_id를 처음 등장한 순서대로 반환한다.
// 같은 작업에 "ok"가 한 번이라도 나오면 완료로 보고, 그 뒤 로그는 무시한다.

func Incomplete(logs []Log) []int {
	done := map[int]bool{}
	seen := map[int]bool{}
	order := []int{}

	for _, l := range logs {
		if !seen[l.JobID] { // 처음 보는 작업이면
			seen[l.JobID] = true
			order = append(order, l.JobID)
		}
		if l.Status == "ok" { // ok면
			done[l.JobID] = true
		}
	}

	result := []int{}
	for _, id := range order {
		if !done[id] { // 끝나지 않은 작업이면
			result = append(result, id)
		}
	}
	return result
}
