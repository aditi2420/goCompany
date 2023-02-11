BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOINSTALL=$(GO) install
GOLIST=$(GO) list
GOVET=$(GO) vet
GOFMT=$(shell which gofmt)
GOSEC=$(shell which gosec)
GOLINT=$(shell which golint)

unexport GOPATH
export GO111MODULE := on
export GOBIN := $(BUILDPATH)/bin

build-server:
	@echo "start building server..."
	@cd $(BUILDPATH)/src/go-company && $(GOINSTALL)
	@echo "DONE!!!"

clean:
	@echo "start cleaning..."
	@rm -f $(BUILDPATH)/bin/go-company
	@echo "DONE!!!"

vet vet-server:
	@echo "start vet on server..."
	@cd $(BUILDPATH)/src/go-company && $(GOVET) `$(GOLIST)` 2>&1
	@echo "DONE!!!	"

lint lint-server:
	@echo "start lint on server..."
	@cd $(BUILDPATH)/src/go-company && $(GOLINT) `$(GOLIST)` 2>&1
	@echo "DONE!!!"

