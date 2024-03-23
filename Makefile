PROJECT := DCUpdate
VERSION := $(shell git describe --tag --abbrev=0)
NEXT_VERSION:=$(shell git describe --tags --abbrev=0 | awk -F . '{OFS="."; $$NF+=1; print}')
SHA1 := $(shell git rev-parse HEAD)
NOW := $(shell date -u +'%Y%m%d-%H%M%S')
TEMPLATE_PREFIX := $(shell cat .env | grep TEMPLATE_PREFIX | awk -F '=' '{print $$2}')

build:
	go generate
	go build -o bin/$(PROJECT) -ldflags "-X main.GITCOMMIT=$(SHA1)" -ldflags "-X main.VERSION=$(VERSION) -X main.BUILDTIME=$(NOW)"

release: fmt
	@git tag -a $(NEXT_VERSION) -m "Release $(NEXT_VERSION)"
	@git push --all
	@git push --tags

newfeature:
	@git checkout -b "Feature-$(NOW)"

reset:
	@git checkout master
	@git pull
	@git fetch
	@git reset --hard origin/master

fmt:
	@go mod tidy
	@goimports -w .
	@gofmt -w -s .
	@go clean ./...


commit: fmt
	@git add .
	@git commit -a -m "$(m)"
	@git pull
	@git push

test:
	@go tool cover -func profile.cov

watch: fmt
	@gow -c run .

run:
	@go run .

