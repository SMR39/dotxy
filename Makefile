package = github.com/dotxy
TAG := $(shell git tag | sort -r | head -n 1)

.PHONY: install release test

release: deps
	goreleaser --rm-dist

deps:
	glide install

install:
	cp dist/darwin-amd64/dotxy /usr/local/bin/dotxy
	chmod +x /usr/local/bin/dotxy

docs:
	sh ./generate_docs.sh

CONTAINER_NAME="dotxy"
build:
	docker build -t $(CONTAINER_NAME) .

# test:
# 	go test -v -cover `go list ./... | grep -v vendor` | sed ''/PASS/s//`printf "\033[32mPASS\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/''
