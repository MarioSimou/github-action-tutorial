unit:
	go test -cover ./internal/...
fmt:
	go fmt ./...
vet:
	go vet ./...