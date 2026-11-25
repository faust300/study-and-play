// Package drills — 002. map·slice 몸풀기
//
// map과 slice를 다루는 감각을 되살리는 짧은 문제 5개. 각 함수 위 주석에 문제와 예시가 있다.
//
//	1. CountWords    단어별 등장 횟수 세기
//	2. Dedupe        중복 제거 (처음 등장한 순서 유지)
//	3. GroupByLength 문자열을 길이별로 묶기
//	4. TwoSum        합이 target인 두 원소의 인덱스 찾기 (O(N))
//	5. Reverse       slice를 제자리에서 뒤집기
//
// 연습하는 Go 문법: make, m[k]++, v, ok := m[k], map[int][]string, 다중 반환,
// a[i], a[j] = a[j], a[i]
package drills

// 1. CountWords: 단어별 등장 횟수를 센다.
// 예) ["a","b","a"] → map{"a":2, "b":1}
func CountWords(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts;
}

// 2. Dedupe: 중복을 제거하되 처음 등장한 순서를 유지한다.
// 예) [3,1,3,2,1] → [3,1,2]
func Dedupe(nums []int) []int {
	seen := map[int]bool{}
	w := 0
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			nums[w] = v
			w++
		}
	}
	return nums[:w]
}

// 3. GroupByLength: 문자열을 길이별로 묶는다. 각 묶음 안의 순서는 입력 순서를 따른다.
// 예) ["go","ts","rust","js"] → map{2:["go","ts","js"], 4:["rust"]}
func GroupByLength(words []string) map[int][]string {
	result := map[int][]string{}

	for _, word := range words {
		l := len(word)
		result[l] = append(result[l], word)
	}

	return result
}

// 4. TwoSum: 합이 target인 두 원소의 인덱스(i < j)를 반환한다. 없으면 ok=false.
// 답이 여러 개면 j가 가장 작은 쌍을 반환한다. O(N)으로 풀어 보자.
// 예) [2,7,11,15], 9 → (0, 1, true)
func TwoSum(nums []int, target int) (int, int, bool) {
	idx := make(map[int]int) // 값 → 인덱스
	for j, v := range nums {
		need := target - v
		if i, ok := idx[need]; ok {
			return i, j, true
		}
		idx[v] = j                 // 나를 기록
	}
	return 0, 0, false
}

// 5. Reverse: 새 슬라이스를 만들지 말고 입력 슬라이스를 제자리에서 뒤집는다.
// 예) [1,2,3,4] → [4,3,2,1]
func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
