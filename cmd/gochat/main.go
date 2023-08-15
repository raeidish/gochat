package main

import (
	"log"
    "github.com/raeidish/gochat/internal/server"
)

func main() {
    err := server.StartServer(":8080") 
    if err != nil {
        log.Fatal(err.Error())
    }
}
