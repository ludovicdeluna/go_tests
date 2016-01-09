export GO15VENDOREXPERIMENT=1

default: install

install:
	@echo "install package in $(GOPATH)/bin"
	@go install github.com/ludovicdeluna/go_tests/...

run: install
	@$(GOPATH)/bin/play

list:
	ls -l $(GOPATH)/bin

version:
	clear
	@echo "-- Current Folder $(notdir $(GOPATH)) --"
	@echo "GOPATH: $$GOPATH"
	@echo "src   : $(GOPATH)/src"
	@echo "pkg   : $(GOPATH)/pkg"
	@echo "bin   : $(GOPATH)/bin"
	@echo "\n-- Runtime --"
	@echo "Required Go 1.5+"
	@echo "Vendor enabled : GO15VENDOREXPERIMENT=$$GO15VENDOREXPERIMENT"
	@go version
	@echo
