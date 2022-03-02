.PHONY: pre-push go-mod check-commit
pre-push: go-mod check-commit

go-mod:
	@echo "mod tidy start..."
	@go mod tidy -v
	@go mod download -x
	@echo "mod tidy complete."

check-commit:
	@echo "testing start..."
	@go clean -testcache
	go test -v ./...
	@echo "test complete."
	@echo "linting start..."
	golangci-lint run -v ./...
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
