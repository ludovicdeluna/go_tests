#mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
#current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
root_dir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
ROOTPATH=$(root_dir)
#Before Go1.5, I prefer keep a vendor path.
# GOPATH:=$(ROOTPATH):$(ROOTPATH)/vendor
GOPATH:=$(ROOTPATH)
#But now, it's really more useful. Always use vendor feature since Go1.5 :
export GO15VENDOREXPERIMENT=1

all: install

version:
	@export GOPATH
	@echo "GOPATH = $$GOPATH\n"
	@go version

install: version
	go install hello pointref

test: version
	go test ./...

clean:
	rm -rf $(ROOTPATH)/pkg

reset: clean
	rm -rf $(ROOTPATH)/bin

run: install
	@echo '-----------'
	@echo 'run hello :'
	@$(ROOTPATH)/bin/hello $(ARGS)
