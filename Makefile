VERSION=0.1.1

.PHONY: all
all:
	@echo goose makefile $(VERSION)
	@echo \"make build\" to build

.PHONY: build
build:
	goxc -pv="$(VERSION)" -build-ldflags='-X main.version=$(VERSION)' xc

.PHONY: package
package: man
	goxc -pv="$(VERSION)" -build-ldflags='-X main.version=$(VERSION)'

.PHONY: html
html:
	pushd docs && make html && popd

.PHONY: man
man:
	pushd docs && make man && popd

.PHONY: upload
upload:
	goxc -pv="$(VERSION)" bintray
