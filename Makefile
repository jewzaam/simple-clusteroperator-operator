unexport GOFLAGS

COMMIT_NUMBER=$(shell git rev-list `git rev-list --parents HEAD | egrep "^[a-f0-9]{40}$$"`..HEAD --count)
CURRENT_COMMIT=$(shell git rev-parse --short=7 HEAD)
OPERATOR_VERSION=$(VERSION_MAJOR).$(VERSION_MINOR).$(COMMIT_NUMBER)-$(CURRENT_COMMIT)

OPERATOR_NAME?=$(shell sed -n 's/.*OperatorName .*"\([^"]*\)".*/\1/p' pkg/config/config.go)
OPERATOR_NAMESPACE?=$(shell sed -n 's/.*OperatorNamespace .*"\([^"]*\)".*/\1/p' pkg/config/config.go)

BINFILE=build/_output/bin/$(OPERATOR_NAME)
MAINPACKAGE=./cmd/manager
GOENV=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GOBUILDFLAGS=-gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}"
TESTOPTS:=

default: gobuild

.PHONY: gobuild
gobuild: gocheck gotest ## Build binary
	${GOENV} go build ${GOBUILDFLAGS} -o ${BINFILE} ${MAINPACKAGE}

.PHONY: gocheck
gocheck: ## Lint code
	gofmt -s -w $$(go list -f '{{ .Dir }}' ./... )
	go vet ./cmd/... ./pkg/...

.PHONY: gotest
gotest:
	go test $(TESTOPTS) $$(go list -mod=readonly -e ./...)
