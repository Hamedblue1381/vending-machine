app:
	@go build -o bin/vending-machine
	@./bin/vending-machine
.PHONY: app