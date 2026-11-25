# alg

문제마다 폴더 하나, `solution.go`에 풀이를 쓰고 테스트로 채점한다.

```bash
go test ./001-callback-retry/          # 채점
go test -v ./001-callback-retry/       # 케이스별 결과
go test -bench=. ./001-callback-retry/ # N=10^6 성능 확인
go test ./...                          # 전체 문제
```
