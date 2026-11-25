// Package parse — 004. 파싱·에러·메서드
//
// 001의 콜백 로그를 이번에는 텍스트 한 줄씩 받아서 직접 파싱한다.
// 각 함수 위 주석에 문제와 예시가 있다.
//
//  1. ParseLine  "3,ok" 같은 한 줄을 Log로 바꾸고, 잘못된 줄은 에러를 반환
//  2. Queue      재시도 대기열. 메서드(Push, Pop, Len)를 가진 타입 만들기
//  3. Summarize  여러 줄을 파싱해 미완료 작업과 잘못된 줄 번호를 함께 반환
//
// 연습하는 Go 문법: error 반환, errors.New, fmt.Errorf("%w"), errors.Is,
// strings.Split·TrimSpace, strconv.Atoi, 메서드와 포인터 리시버(*Queue)
package parse

import "errors"

// Log는 콜백 로그 한 줄이다. Status는 "fail" 또는 "ok".
type Log struct {
	JobID  int
	Status string
}

// ParseLine이 반환하는 에러. 호출하는 쪽은 errors.Is로 종류를 구분한다.
var (
	ErrFormat = errors.New("형식 오류")     // 쉼표로 나눈 칸이 2개가 아님
	ErrJobID  = errors.New("job_id 오류") // job_id가 정수가 아님
	ErrStatus = errors.New("status 오류") // status가 "ok"나 "fail"이 아님
)

// 1. ParseLine: "job_id,status" 형식의 한 줄을 Log로 바꾼다.
// 각 칸 앞뒤의 공백은 무시한다. 잘못된 줄이면 위의 에러 중 하나를 반환한다.
// 에러를 감싸서(fmt.Errorf + %w) 어떤 줄이 왜 틀렸는지 메시지를 덧붙여도 된다.
// 예) "3,ok"        → Log{3, "ok"}, nil
//
//	" 7 , fail "  → Log{7, "fail"}, nil
//	"3"           → ErrFormat
//	"abc,ok"      → ErrJobID
//	"3,done"      → ErrStatus
func ParseLine(line string) (Log, error) {
	// TODO
	return Log{}, nil
}

// 2. Queue: 재시도할 job_id를 넣고 먼저 넣은 것부터 꺼내는 대기열(FIFO).
// 아무 초기화 없이 `var q Queue`로 바로 쓸 수 있어야 한다.
//
//	q.Push(3); q.Push(1)
//	q.Pop() → 3, true
//	q.Len() → 1
//	빈 대기열에서 q.Pop() → 0, false
type Queue struct {
	// TODO: 필요한 필드를 추가한다.
}

func (q *Queue) Push(jobID int) {
	// TODO
}

func (q *Queue) Pop() (int, bool) {
	// TODO
	return 0, false
}

func (q *Queue) Len() int {
	// TODO
	return 0
}

// 3. Summarize: 여러 줄을 파싱해서 두 가지를 반환한다.
//   - incomplete: 완료되지 못한 job_id (001과 같은 규칙, 처음 등장한 순서)
//   - badLines:   파싱에 실패한 줄의 번호 (0부터 시작, 오름차순)
//
// 잘못된 줄은 건너뛰고 나머지 줄로 계속 처리한다.
// 예) ["3,fail", "x,ok", "1,fail", "3,ok", "2"] → incomplete [1], badLines [1 4]
func Summarize(lines []string) (incomplete []int, badLines []int) {
	// TODO
	return nil, nil
}
