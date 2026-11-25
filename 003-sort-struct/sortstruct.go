// Package sortstruct — 003. 정렬·구조체
//
// 구조체 slice를 정렬하고 가공하는 문제 3개. 각 함수 위 주석에 문제와 예시가 있다.
//
//  1. RankMentors  정산 금액 내림차순, 같으면 이름 오름차순으로 정렬 (입력은 바꾸지 않음)
//  2. TopK         가장 많이 나온 단어 k개 (횟수 같으면 단어 오름차순)
//  3. MergeWindows 겹치거나 맞닿은 점검 시간대 [Start, End)를 합치기
//
// 연습하는 Go 문법: 구조체, sort.Slice, slice 복사 (append([]T(nil), s...))
package sortstruct

import (
	"sort"
)

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
	wordCount := make(map[string]int)
	result := make([]string, 0, k)
	for _, word := range words {
		wordCount[word]++
		if wordCount[word] == 1 {
			result = append(result, word)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if wordCount[a] != wordCount[b] {
			return wordCount[a] > wordCount[b]   // 횟수가 많은 쪽이 앞
		}
		return a < b       // 같으면 사전순으로 앞서는 쪽이 앞
	})
	if len(result) > k {
		result = result[:k]
	}
	return result
}

// 1. slices.SortFunc + cmp.Compare: sort.Slice는 인덱스 i, j를 받아서 result[i]를 꺼내야 했어요. SortFunc는 값을 바로 받고, TS의 sort((a, b) => ...)처럼 음수, 0, 양수를 돌려줘요. TS에서 쓰던 방식과 더 비슷하죠.
// 2. cmp.Or: 앞의 비교가 0(같음)이면 다음 비교를 써요. "횟수로 비교하고, 같으면 이름으로"를 if 없이 한 줄로 쓸 수 있어요. 내림차순은 count[b], count[a]처럼 인자 순서를 뒤집으면 돼요.
// 3. min(k, len(result)): Go 1.21부터 들어온 내장 함수예요. if len(result) > k 블록이 한 줄로 줄어요.
// 4. if count[w]++; count[w] == 1: if 안에 짧은 문장을 먼저 실행할 수 있어요. 앞에서 쓴 if i, ok := idx[need]; ok와 같은 문법이에요. 다만 이건 오히려 읽기 어려워질 수 있어서, 원래처럼 두 줄로 써도 충분해요.
// func TopK(words []string, k int) []string {
// 	count := map[string]int{}
// 	result := []string{}
// 	for _, w := range words {
// 		if count[w]++; count[w] == 1 {
// 			result = append(result, w)
// 		}
// 	}
// 	slices.SortFunc(result, func(a, b string) int {
// 		return cmp.Or(cmp.Compare(count[b], count[a]), cmp.Compare(a, b))
// 	})
// 	return result[:min(k, len(result))]
// }

// Window는 점검 시간대다. Start 이상, End 미만.
type Window struct {
	Start int
	End   int
}

// 3. MergeWindows: 겹치거나 맞닿은 시간대를 합쳐 시작 시각 순으로 반환한다.
// 입력은 정렬돼 있지 않을 수 있다. 맞닿음: [1,3)과 [3,5)는 [1,5)로 합친다.
// 예) [{8 10} {1 3} {2 6} {6 7}] → [{1 7} {8 10}]
func MergeWindows(windows []Window) []Window {
	sorted := append([]Window(nil), windows...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Start < sorted[j].Start
	})
	result := []Window{}
	for _, w := range sorted {
		n := len(result)
		if n > 0 && w.Start <= result[n-1].End {
			result[n-1].End = max(result[n-1].End, w.End)
		} else {
			result = append(result, w)
		}
	}
	return result
}
