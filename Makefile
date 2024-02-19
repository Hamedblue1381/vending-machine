app:
	@go build -o bin/cli/vending-machine
	@./bin/cli/vending-machine
.PHONY: app