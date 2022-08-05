GO        	?= GO111MODULE=on go
STATIC_PATH	= static/cert


.PHONY: bin

bin:
	$(GO) build -o bin/app cmd/main.go

run:
	$(GO) run cmd/main.go

pb:
	rm -rf ./generate/*
	protoc --go_out=./generate --go_opt=paths=source_relative \
		--go-grpc_out=./generate --go-grpc_opt=paths=source_relative \
		proto/user.proto proto/svc.proto

tidy:
	$(GO) mod tidy

cert:
	mkdir -p $STATIC_PATH
	# openssl req -new -x509 -keyout static/cert/server.key -out static/cert/server.crt
	openssl genrsa -out "${STATIC_PATH}/server.key" 2048
	openssl req -new -key "${STATIC_PATH}/server.key" -out "${STATIC_PATH}/server.csr"
	openssl x509 -req -days 365 -in "${STATIC_PATH}/server.csr" -signkey "${STATIC_PATH}/server.key" -out "${STATIC_PATH}/server.crt"

downloadproto:
	sudo chmod +x download_proto.sh
	./download_proto.sh