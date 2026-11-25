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
