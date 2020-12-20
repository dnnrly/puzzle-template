
.PHONY: run-all
run-all:
	go run cmd/euler/main.go

.PHONY: latest
latest:
	go run cmd/euler/main.go latest

.PHONY: next
next:
	go run cmd/next/main.go