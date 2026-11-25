package sortstruct
import "sort"
// Mentor는 멘토 한 명의 정산 정보다.
type Mentor struct {
	Name   string
	Amount int // 정산 금액
}

// 1. RankMentors: 정산 금액이 큰 순서로 정렬한 새 슬라이스를 반환한다.
// 금액이 같으면 이름 오름차순. 입력 슬라이스는 바꾸지 않는다.
// 예) [{b 100} {a 300} {c 100}] → [{a 300} {b 100} {c 100}]
func RankMentors(mentors []Mentor) []Mentor {
	sorted := make([]Mentor, len(mentors))
	copy(sorted, mentors)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Amount == sorted[j].Amount {
			return sorted[i].Name < sorted[j].Name
		}
		return sorted[i].Amount > sorted[j].Amount
	})
	return sorted
}

// 2. TopK: 가장 많이 나온 단어 k개를 많이 나온 순서로 반환한다.
// 횟수가 같으면 단어 오름차순. 단어 종류가 k개보다 적으면 있는 만큼만.
// 예) ["go","ts","go","rust","ts","go"], k=2 → ["go","ts"]
func TopK(words []string, k int) []string {
	// TODO
	return nil
}

// Window는 점검 시간대다. Start 이상, End 미만.
type Window struct {
	Start int
	End   int
}

// 3. MergeWindows: 겹치거나 맞닿은 시간대를 합쳐 시작 시각 순으로 반환한다.
// 입력은 정렬돼 있지 않을 수 있다. 맞닿음: [1,3)과 [3,5)는 [1,5)로 합친다.
// 예) [{8 10} {1 3} {2 6} {6 7}] → [{1 7} {8 10}]
func MergeWindows(windows []Window) []Window {
	// TODO
	return nil
}
