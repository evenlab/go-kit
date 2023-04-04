.PHONY: pre-push
pre-push: go-mod-tidy check-commit

go-mod-tidy:
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
	@go mod tidy
	@go mod download
	@echo "mod tidy complete."

.PHONY: go-mod-upgrade
go-mod-upgrade: go-mod-update go-mod-tidy

go-mod-update:
	@echo "mod upgrade start..."
	@cd ./base58 && go get -u ./...
	@cd ./bytes && go get -u ./...
	@cd ./context && go get -u ./...
	@cd ./crypto && go get -u ./...
	@cd ./drive && go get -u ./...
	@cd ./equal && go get -u ./...
	@cd ./errors && go get -u ./...
	@cd ./merkle && go get -u ./...
	@cd ./network && go get -u ./...
	@cd ./strings && go get -u ./...
	@cd ./timestamp && go get -u ./...
	@cd ./zero && go get -u ./...
	@go get -u ./...
	@echo "mod upgrade complete."

.PHONY: check-commit
check-commit: tests lints

tests:
	@echo "testing start..."
	@go clean -testcache
	@cd ./base58 && go test -race -short ./...
	@cd ./bytes && go test -race -short ./...
	@cd ./context && go test -race -short ./...
	@cd ./crypto && go test -race -short ./...
	@cd ./drive && go test -race -short ./...
	@cd ./equal && go test -race -short ./...
	@cd ./errors && go test -race -short ./...
	@cd ./merkle && go test -race -short ./...
	@cd ./network && go test -race -short ./...
	@cd ./strings && go test -race -short ./...
	@cd ./timestamp && go test -race -short ./...
	@cd ./zero && go test -race -short ./...
	@echo "test complete."

lints:
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

.PHONY: protoc-upgrade
protoc-upgrade: protoc-update generate-proto go-generate-proto

protoc-update:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: go-generate-proto
go-generate-proto: generate-proto go-generate-after

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

go-generate-after:
	@go install github.com/favadi/protoc-go-inject-tag@latest && \
	protoc-go-inject-tag -input=crypto/proto/pb/*.pb.go && \
	protoc-go-inject-tag -input=timestamp/proto/pb/*.pb.go
	@go mod tidy
