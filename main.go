package main

import (
	"fmt"
	"log"

	"github.com/prongbang/wire-example/pkg/greeter"
)

func main() {
	g, err := greeter.Initialize()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
	}
	log.Println(g.GetAll())
}
