unexport GOFLAGS
GOENV=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GOBUILDFLAGS=-gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}"
MAINPACKAGE=./cmd/manager
TESTOPTS:=

default: gobuild

.PHONY: gobuild
gobuild: gocheck gotest ## Build binary
	${GOENV} go build ${GOBUILDFLAGS} -o ${BINFILE} ${MAINPACKAGE}

.PHONY: gocheck
gocheck: ## Lint code
	gofmt -s -l $$(go list -f '{{ .Dir }}' ./... ) | grep ".*\.go"; if [ "$$?" = "0" ]; then gofmt -s -d $$(go list -f '{{ .Dir }}' ./... ); exit 1; fi
	go vet ./cmd/... ./pkg/...

.PHONY: gotest
gotest:
	go test $(TESTOPTS) $$(go list -mod=readonly -e ./...)
