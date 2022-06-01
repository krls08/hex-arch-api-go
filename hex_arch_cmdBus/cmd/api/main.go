package main

import (
	"log"

	"github.com/krls08/hex-arch-api-go/hex_arch/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
