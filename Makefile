#mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
#current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
ROOTPATH=$(root_dir)
GOPATH:=$(ROOTPATH):$(ROOTPATH)/vendor

all: install

version:
	@export GOPATH
	@echo "GOPATH = $$GOPATH\n"
	@go version

install: version
	go install hello

test: version
	go test ./...

clean:
	rm -rf $(ROOTPATH)/pkg
	rm -rf $(ROOTPATH)/vendor/pkg

reset: clean
	rm -rf $(ROOTPATH)/bin

run: install
	@echo '-----------'
	@echo 'run hello :'
	@$(ROOTPATH)/bin/hello $(ARGS)
