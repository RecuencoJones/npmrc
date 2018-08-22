.phony: default
default: x-compile

# Install dev and runtime dependencies
.phony: setup
setup:
	go get github.com/mitchellh/gox
	go get -t ./...

# Assemble local distribution binary
.phony: build
build: setup
	go build -o build/npmrc ./...

# Install application to go binaries
.phony: install
install: setup
	go install ./...

# Cross compile to generate binaries and .tar.gz for all dists
.phony: x-compile
x-compile: setup
	gox -output="bin/{{.OS}}_{{.Arch}}/{{.Dir}}" ./...
	for i in bin/*; do tar -zvcf "$$i.tar.gz" "$$i"; done
