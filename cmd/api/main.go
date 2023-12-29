package main

import (
	"log"
	"github.com/marcohb99/go-api-example/cmd/api/bootstrap"
)

func main() {
	// encapsulate run function if we want to test it aside
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}