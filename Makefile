DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
# HTTPS_GIT := https://github.com/scalarorg/scalar-core.git

proto-all: proto-update-deps proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@DOCKER_BUILDKIT=1 docker build -t scalar/proto-gen -f ./Dockerfile .
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace scalar/proto-gen sh ./scripts/protocgen.sh
	@echo "Generating Protobuf Swagger endpoint"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace scalar/proto-gen sh ./scripts/protoc-swagger-gen.sh
	@statik -src=./client/docs/static -dest=./client/docs -f -m

proto-format:
	@echo "Formatting Protobuf files"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace tendermintdev/docker-build-proto \
	$( find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; )

proto-lint:
	@echo "Linting Protobuf files"
	@$(DOCKER_BUF) lint

# proto-check-breaking:
# 	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=main

proto-update-deps:
	@echo "Updating Protobuf deps"
	@$(DOCKER_BUF) mod update

.PHONY: proto-all proto-gen proto-gen-any proto-format proto-lint proto-check-breaking proto-update-deps


# Install all generate prerequisites
.Phony: prereqs
prereqs:
	@which mdformat &>/dev/null || ( \
		echo "Installing mdformat in a virtual environment..." && \
		python3 -m venv .venv && \
		. .venv/bin/activate && \
		pip install mdformat )
	@which protoc &>/dev/null || echo "Please install protoc for grpc (https://grpc.io/docs/languages/go/quickstart/)"
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/matryer/moq@latest
	go install github.com/rakyll/statik@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
