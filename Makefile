.PHONY: bin

bin:
	go build -o bin/app cmd/main.go

run:
	go run cmd/main.go

pb:
	rm -rf ./generate/*
	protoc --go_out=./generate --go_opt=paths=source_relative \
		--go-grpc_out=./generate --go-grpc_opt=paths=source_relative \
		proto/user.proto proto/svc.proto

tidy:
	go mod tidy