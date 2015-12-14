SRC_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR:=$(shell dirname $(SRC_DIR))
GOPATH:=$(ROOT_DIR)
export GO15VENDOREXPERIMENT=1

all: install clean end_msg

run: run_hello end_msg

test: test_all end_msg

check_src:
ifneq ("$(notdir $(SRC_DIR))","src")
	$(error Move content of this folder into $(notdir $(SRC_DIR))/src)
endif

version: check_src
	clear
	@echo "-- Projet $(notdir $(ROOT_DIR)) --"
	@echo "GOPATH: $$GOPATH"
	@echo "src   : $(SRC_DIR)"
	@echo "pkg   : $(ROOT_DIR)/pkg"
	@echo "bin   : $(ROOT_DIR)/bin"
	@echo "\n-- Runtime --"
	@echo "Required Go 1.5+"
	@echo "Vendor enabled : GO15VENDOREXPERIMENT=$$GO15VENDOREXPERIMENT"
	@go version
	@echo

install: version
	@echo "-- Build and Install --"
	go install hello pointref
	@echo

test_all: version
	go test ./...

clean_msg: check_src
	@echo "-- Clean temporary objects --"

clean: clean_msg
	@echo "pkg objects"
	@rm -rf $(ROOT_DIR)/pkg

reset: clean
	@echo "binary"
	@rm -rf $(ROOT_DIR)/bin

end_msg:
	@echo "\nTerminated."

run_hello: install
	@echo "-- Run command hello (hello.go) --"
	@$(ROOT_DIR)/bin/hello $(ARGS)
