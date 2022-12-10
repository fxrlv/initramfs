GOFLAGS ?= -trimpath
LDFLAGS ?= -s -w


.PHONY: build
build:
	CGO_ENABLED=0 go build $(GOFLAGS) \
		-ldflags '$(LDFLAGS)' \
		-o build/init \
		./cmd/init

.PHONY: fmt
fmt:
	@gofmt -l -s -w .

.PHONY: kernel
kernel:
	@docker buildx build --target kernel -o build .

.PHONY: initramfs
initramfs:
	@docker buildx build --target initramfs -o build .
