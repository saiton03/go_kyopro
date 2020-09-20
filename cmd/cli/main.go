package main

import (
	"go_kyopro/assemble"
	"log"
	"os"
)

func main() {
	if err := assemble.RootCmd.Execute(); err != nil {
		log.Fatalf("%s: %v", os.Args[0], err)
	}
}
