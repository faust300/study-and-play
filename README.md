# study-and-play

Go 학습용 알고리즘 문제 저장소. Claude가 문제를 내고 내가 Go로 푼다.

알고리즘 자체보다 **Go 문법과 표준 라이브러리에 익숙해지는 것**이 목적이다. 그래서 문제마다 새로 써 볼 Go 문법(`:=`, `range`, map, slice, 다중 반환, `sort.Slice` 등)이 하나씩 들어가도록 순서를 잡는다.

## 진행 방식

1. **Claude가 문제를 낸다.** 문제 설명과 빈 함수(`// TODO`), 채점용 테스트를 함께 만든다. 테스트는 Claude가 정답 풀이로 미리 검증해 둔다(정답 코드는 저장소에 넣지 않는다).
2. **내가 푼다.** 함수 안을 채우고 `go test`로 채점한다.
3. **막히면 Claude에게 묻는다.** 정답 대신 어디서 틀렸는지, 다음에 무엇을 생각할지 힌트를 받는다.

문제 주제는 실무에서 겪은 상황(콜백 재시도, 정산, 점검 시간대 등)에서 가져온다.

## 문제 목록

| 폴더 | 내용 | 새로 써 보는 Go 문법 | 상태 |
| --- | --- | --- | --- |
| `001-callback-retry` | 콜백 재시도 로그에서 미완료 작업 찾기 | `:=`, `for`·`range`, `_`, `map[int]bool`, `append` | 완료 |
| `002-map-slice-drills` | map·slice 몸풀기 5문제 (세기, 중복 제거, 묶기, TwoSum, 뒤집기) | `make`, `m[k]++`, `v, ok := m[k]`, `map[int][]string`, 다중 반환, 값 맞바꾸기 | 완료 |
| `003-sort-struct` | 정렬·구조체 3문제 (정산 순위, Top K, 점검 시간대 병합) | 구조체, `sort.Slice`, slice 복사 | 진행 전 |

## 실행

```bash
go test ./...                              # 전체 채점
go test -v ./002-map-slice-drills/         # 문제 묶음별 케이스 결과
go test -v -run TestTwoSum ./002-map-slice-drills/  # 한 문제만
go test -bench=. ./001-callback-retry/     # 성능 확인 (N=10^6)
```

Go 1.23 이상이 필요하다.
