PROGRAM = askamanpage

.PHONY: all
all: $(PROGRAM)

#
# env
#

CMD_CLANG ?= clang
CMD_GO ?= go
CMD_GREP ?= grep
CMD_RM ?= rm

#
# make
#

.ONESHELL:
SHELL = /bin/bash

PARALLEL = $(shell $(CMD_GREP) -c ^processor /proc/cpuinfo)
MAKE = make
MAKEFLAGS += --no-print-directory

#
# environment
#

DEBUG ?= 0
UNAME_M := $(shell uname -m)
UNAME_R := $(shell uname -r)

ifeq ($(DEBUG),1)
	GO_DEBUG_FLAG =
else
	GO_DEBUG_FLAG = -w
endif

ifeq ($(UNAME_M),x86_64)
	ARCH = x86_64
	LINUX_ARCH = x86
	GO_ARCH = amd64
endif

ifeq ($(UNAME_M),aarch64)
	ARCH = arm64
	LINUX_ARCH = arm64
	GO_ARCH = arm64
endif

#
# common variables
#

STATIC ?= 0

GO_TAGS =

ifeq ($(STATIC), 1)
    GO_TAGS := $(GO_TAGS),netgo
endif

GO_ENV =
GO_ENV += GOOS=linux
GO_ENV += CC=$(CMD_CLANG)
GO_ENV += GOARCH=$(GO_ARCH)

#
# program
#

SRC_DIRS = ./cmd/ ./pkg/
SRC = $(shell find $(SRC_DIRS) -type f -name '*.go' ! -name '*_test.go')

$(PROGRAM): $(SRC)
	$(GO_ENV) $(CMD_GO) build \
		-tags $(GO_TAGS) \
		-ldflags="$(GO_DEBUG_FLAG)" \
		-v -o $@ \
		./cmd/$(PROGRAM)

#
# clean
#

.PHONY: clean
clean:
#
	$(CMD_RM) -f $(PROGRAM)
