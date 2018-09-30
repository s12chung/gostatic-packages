test:
	go test ./...

test-report:
	go test -v -covermode=atomic -coverprofile=coverage.out ./...

lint:
	golangci-lint run ./...