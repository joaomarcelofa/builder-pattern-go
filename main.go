package main

import (
	"fmt"
	"log"

	"github.com/joaomarcelofa/builder_factory/mac/factory"
)

func main() {
	macBookAir, err := factory.CreateNewMac(factory.MacBookAir)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%+v", macBookAir)
}
