TARGETOS = linux
APP_NAME = core
APP_VERSION = `git rev-parse --short HEAD | xargs git describe`
GOLDFLAGS = "-X core/internal.version=$(APP_VERSION)"
CONTAINER_IMG_NAME= $(APP_NAME):$(APP_VERSION)
CONTAINER_IMG_NAME_BUILDER= $(APP_NAME)-buider:$(APP_VERSION)
TZ = Asia/Yangon
BIN_ARC_FILE_NAME = $(APP_NAME)-$(APP_VERSION)-$(TARGETOS).tar.gz



.PHONY: check
check:
	@echo $(APP_VERSION)
	@echo $(GOLDFLAGS)
	


.PHONY: version-git
version-git:
	@echo $(APP_VERSION)
	@echo $(GOLDFLAGS)

.PHONY: format
format:
	go fmt ./...

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: test
test: format mod-tidy
	go test ./...
	go run local/main.go

.PHONY: build
build: format mod-tidy 
	GOOS=$(TARGETOS) CGO_ENABLED=0 go build -ldflags $(GOLDFLAGS)


	




