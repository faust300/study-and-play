# study-and-play

Go 학습용 알고리즘 문제 저장소. 문제와 테스트는 Claude가 출제하고, 풀이는 직접 작성한다.

알고리즘 자체보다 **Go 문법과 표준 라이브러리에 익숙해지는 것**이 목적이다. 그래서 문제마다 새로 써 볼 Go 문법(`:=`, `range`, map, slice, 다중 반환, `sort.Slice` 등)이 하나씩 들어가도록 순서를 잡는다.

## 진행 방식

1. **출제** — Claude가 문제 설명, 빈 함수(`// TODO`), 채점용 테스트를 만든다. 테스트는 정답 풀이로 미리 검증해 두고, 정답 코드는 저장소에 넣지 않는다.
2. **풀이** — 빈 함수를 채우고 `go test`로 채점한다.
3. **힌트** — 막히면 Claude에게 정답 대신 틀린 지점과 다음에 생각할 방향을 묻는다.

문제 주제는 실무에서 겪은 상황(콜백 재시도, 정산, 점검 시간대 등)에서 가져온다.

## 문제 목록

| 폴더 | 내용 | 새로 써 보는 Go 문법 | 상태 |
| --- | --- | --- | --- |
| `001-callback-retry` | 콜백 재시도 로그에서 미완료 작업 찾기 | `:=`, `for`·`range`, `_`, `map[int]bool`, `append` | 완료 |
| `002-map-slice-drills` | map·slice 몸풀기 5문제 (세기, 중복 제거, 묶기, TwoSum, 뒤집기) | `make`, `m[k]++`, `v, ok := m[k]`, `map[int][]string`, 다중 반환, 값 맞바꾸기 | 완료 |
| `003-sort-struct` | 정렬·구조체 3문제 (정산 순위, Top K, 점검 시간대 병합) | 구조체, `sort.Slice`, slice 복사, `max`·`min` | 완료 |
| `004-parse-error-method` | 콜백 로그 텍스트 파싱, 재시도 대기열, 요약 | `error` 반환, `errors.Is`, `fmt.Errorf("%w")`, `strings`·`strconv`, 메서드와 포인터 리시버 | 진행 전 |

## 실행

```bash
go test ./...                              # 전체 채점
go test -v ./002-map-slice-drills/         # 문제 묶음별 케이스 결과
go test -v -run TestTwoSum ./002-map-slice-drills/  # 한 문제만
go test -bench=. ./001-callback-retry/     # 성능 확인 (N=10^6)
```

Go 1.23 이상이 필요하다.
