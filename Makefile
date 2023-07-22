

run:
	go mod tidy && go run cmd/exchange/main.go

test:
	go mod tidy && go test ./...
