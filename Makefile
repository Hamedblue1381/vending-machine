app:
	@go build -o bin/vending-machine
	@./bin/vending-machine
test:
	go test -v -cover ./...
build:
	docker-compose up -d --build
cli:
	@docker-compose build
	@docker-compose run --service-ports --rm app -e CLI_ENABLED=true app
.PHONY: app test cli