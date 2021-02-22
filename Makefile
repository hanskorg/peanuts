GOOS       = linux
GOARCH     = amd64
NAME       = peanuts
BIN_PATH   = $(shell pwd)/bin
VERSION?   =

imageName = $(NAME)

ifneq ($(ver),)
    VERPERFIX = -$(ver)
endif

ifneq ($(VERSION),)
    imageName = hansk887/$(NAME):$(VERSION)$(VERPERFIX)
else
    VERSION   = $(shell git rev-parse --short=8 HEAD || echo unknown)
    imageName = hansk887/$(NAME):$(VERSION)$(VERPERFIX)
endif


LDFLAGS=-ldflags "-X 'peanuts/main.Version=$(VERSION)$(VERPERFIX)' -X 'peanuts/main.Build=`TZ=UTC-8 date +%FT%T%z`' -X peanuts/main.Name=peanuts"

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build ${LDFLAGS} -o $(BIN_PATH)/api api.go

image: build
	docker build --no-cache . --tag  ${imageName} && \
    docker push $(imageName) && \
    docker rmi -f ${imageName}

release:
	export VERSION=$(CI_BUILD_TAG) && \
	$(MAKE) image

clean:
	rm -rf $(NAME)
