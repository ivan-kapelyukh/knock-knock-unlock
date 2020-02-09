BUILD := go build
RUN := go run
NPM := npm --prefix web

.PHONY: all build run clean

all: build

build: server web/dist

run: web/dist FORCE
	$(RUN) cmd/server/main.go

server: cmd/server/main.go $(wildcard internal/app/server/*.go) $(wildcard internal/app/server/*/*.go)
	$(BUILD) -o $@ $<

web/dist: web/node_modules FORCE
	$(NPM) run-script build

web/node_modules: web/package.json web/package-lock.json
	$(NPM) install

clean:
	$(RM) server

FORCE:
