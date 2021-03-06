package main

import (
	"log"

	"github.com/amine7536/reverse-scan/cmd"
)

const (
	// Version : app version
	Version = "v0.2.3"
)

func main() {

	if err := cmd.NewRootCmd(Version).Execute(); err != nil {
		log.Fatal(err)
	}
}
