package main

import (
	"log"
)

func assert(condition bool, message string) {
	if !condition {
		log.Fatalf("Assertion failed: %s", message)
	}
}
