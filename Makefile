INTERNAL = $(wildcard ./internal/*/*.go)
GOCHAT = ./cmd/gochat/main.go


.PHONY: run

run-gochat: gochat 
	./gochat

gochat: $(INTERNAL) $(GOCHAT)
	go build -o gochat ./cmd/gochat/main.go
