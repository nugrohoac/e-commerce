package main

import (
	"log"

	"github.com/nugrohoac/e-commerce/cmd"
)

func main() {
	if err := cmd.RootCMD.Execute(); err != nil {
		log.Fatalf("Fail init Root CMD, err : %v", err)
	}
}
