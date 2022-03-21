.PHONY: pre-push go-mod check-commit
pre-push: go-mod check-commit

go-mod:
	@echo "mod tidy start..."
	@cd ./base58 && go mod tidy
	@cd ./base58 && go mod download
	@cd ./bytes && go mod tidy
	@cd ./bytes && go mod download
	@cd ./context && go mod tidy
	@cd ./context && go mod download
	@cd ./crypto && go mod tidy
	@cd ./crypto && go mod download
	@cd ./drive && go mod tidy
	@cd ./drive && go mod download
	@cd ./equal && go mod tidy
	@cd ./equal && go mod download
	@cd ./errors && go mod tidy
	@cd ./errors && go mod download
	@cd ./merkle && go mod tidy
	@cd ./merkle && go mod download
	@cd ./network && go mod tidy
	@cd ./network && go mod download
	@cd ./strings && go mod tidy
	@cd ./strings && go mod download
	@cd ./timestamp && go mod tidy
	@cd ./timestamp && go mod download
	@cd ./zero && go mod tidy
	@cd ./zero && go mod download
	@echo "mod tidy complete."

check-commit:
	@echo "testing start..."
	@go clean -testcache
	@cd ./base58 && go test -race ./...
	@cd ./bytes && go test -race ./...
	@cd ./context && go test -race ./...
	@cd ./crypto && go test -race ./...
	@cd ./drive && go test -race ./...
	@cd ./equal && go test -race ./...
	@cd ./errors && go test -race ./...
	@cd ./merkle && go test -race ./...
	@cd ./network && go test -race ./...
	@cd ./strings && go test -race ./...
	@cd ./timestamp && go test -race ./...
	@cd ./zero && go test -race ./...
	@echo "test complete."
	@echo "linting start..."
	@cd ./base58 && golangci-lint run
	@cd ./bytes && golangci-lint run
	@cd ./context && golangci-lint run
	@cd ./crypto && golangci-lint run
	@cd ./drive && golangci-lint run
	@cd ./equal && golangci-lint run
	@cd ./errors && golangci-lint run
	@cd ./merkle && golangci-lint run
	@cd ./network && golangci-lint run
	@cd ./strings && golangci-lint run
	@cd ./timestamp && golangci-lint run
	@cd ./zero && golangci-lint run
	@echo "lint complete."

.PHONY: protoc-upgrade protoc-update generate-proto go-generate-proto
protoc-upgrade: protoc-update generate-proto go-generate-proto

protoc-update:
	go get -u google.golang.org/protobuf@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate-proto:
	@echo "Compiling protobuf files..."
	@rm -f crypto/proto/pb/*
	@rm -f timestamp/proto/pb/*
	# Crypto pb files generation section
	@protoc \
		-I=. \
		--go_out=. \
		--go_opt=module=github.com/evenlab/go-kit \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/evenlab/go-kit \
		--proto_path=crypto/proto crypto/proto/*.proto

	# Timestamp pb files generation section
	@protoc \
		-I=. \
		--go_out=. \
		--go_opt=module=github.com/evenlab/go-kit \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/evenlab/go-kit \
		--proto_path=timestamp/proto timestamp/proto/*.proto
	@echo "Compiling completed."

.PHONY: go-generate-proto generate-proto go-generate-after
go-generate-proto: generate-proto go-generate-after

go-generate-after:
	@go get -d github.com/favadi/protoc-go-inject-tag && \
	protoc-go-inject-tag -input=crypto/proto/pb/*.pb.go && \
	protoc-go-inject-tag -input=timestamp/proto/pb/*.pb.go
