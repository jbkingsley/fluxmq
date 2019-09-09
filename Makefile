# Copyright (c) Drasko DRASKOVIC
# SPDX-License-Identifier: Apache-2.0

PROGRAM = fluxmq
SOURCES = $(wildcard *.go) cmd/main.go

all: $(PROGRAM)

.PHONY: all clean

$(PROGRAM): $(SOURCES)
	go build -ldflags "-s -w" -o $@ cmd/main.go

clean:
	rm -rf $(PROGRAM)
