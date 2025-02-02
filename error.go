package main

import (
	"log"
)

func fatal_check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
