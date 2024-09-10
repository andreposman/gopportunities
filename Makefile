.PHONY: default run build docs clean


#vars
APP_NAME=gopportunities


#tasks
default: run-with-docs

run: 
	@go run cmd/api/main.go

run-with-docs: 
	@swag init -g cmd/api/main.go
	@go run cmd/api/main.go

build:
	@go build -o $(APP_NAME) cmd/api/main.go

test:
	@go test -v ./...

docs:
	@swag init -g cmd/api/main.go

clean:
	@rm -rf $(APP_NAME) 
	@rm -rf docs