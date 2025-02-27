TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=localhost
NAMESPACE=FabioAntunes
NAME=metanetworks
BINARY=terraform-provider-${NAME}
VERSION?=1.0.0-pre-2.4
OS_ARCH?=darwin_amd64

default: install

build:
	go build -o ./bin/${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ./bin/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
