NAME:=$(CI_PROJECT_NAME)
VERSION:=$(CI_COMMIT_REF_NAME)

ifeq ($(VERSION),)
	# Looks like we are not running in the CI so default to current branch
	VERSION:=$(shell git rev-parse --abbrev-ref HEAD)
endif

ifeq ($(NAME),)
	# Looks like we are not running in the CI so default to current directory
	NAME:=$(notdir $(CURDIR))
endif

ifeq ($(VERSION), master)
	VERSION:=latest
endif


.PHONY: build
build: check
	CGO_ENABLED=0 GOOS=linux bash -c 'go build -ldflags "-X main.version=${VERSION}" -o ${NAME} -a -installsuffix cgo main.go'

.PHONY: check
check: lint test

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run --enable gofmt ./...

docker-build: ## Build docker image with short git hash
	@ docker build -t blockdaemon/${NAME}:${VERSION} .

docker-push: ## Upload image to repository
	@ docker push blockdaemon/${NAME}:${VERSION}
