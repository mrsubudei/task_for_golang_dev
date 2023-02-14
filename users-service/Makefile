.PHONY: test
test:
		go test ./...

.PHONY: cover
cover:
		go test ./... -coverprofile=coverage.out
		go tool cover -html=coverage.out