package main

import (
	"fmt"
	"log"

	"github.com/denisbrodbeck/machineid"
)

func main() {
	log.SetFlags(0)
	id, err := machineid.ID()
	if err != nil {
		log.Fatalf("Failed to read machine id with error: %s\n", err)
	}
	fmt.Println(id)
}
