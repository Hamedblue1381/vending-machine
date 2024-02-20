app:
	@go build -o bin/vending-machine
	@./bin/vending-machine
test:
	go test -v -cover ./...

.PHONY: app