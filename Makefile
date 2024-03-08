.PHONY: build run clean

build:
	go build -o auth cmd/main.go

run:
	go run cmd/main.go

depend:
	go mod download && go mod verify

lint:
	golangci-lint run

migrate:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=postgres host=db" up

down:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=postgres host=db" down

reset:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=postgres host=db" reset

clean: reset
	rm auth

gen: ./api/auth/auth.proto
	protoc -I api api/auth/auth.proto --go_out=./pkg/api --go_opt=paths=source_relative --go-grpc_out=./pkg/api --go-grpc_opt=paths=source_relative