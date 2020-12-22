
.PHONY: run-all
run-all:
	go run cmd/puzzle/main.go

.PHONY: latest
latest:
	go run cmd/puzzle/main.go latest

.PHONY: test
test:
	go test .

.PHONY: acceptance-test
acceptance-test:
	go test ./cmd/puzzle
	go test ./cmd/next

.PHONY: next
next:
	go run cmd/next/main.go