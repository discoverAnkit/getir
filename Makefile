.PHONY: tools

tools:
	GO111MODULE=on go generate -tags build tools.go