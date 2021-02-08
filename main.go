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
	fmt.Printf("MacBookAir:\n%+v\n", macBookAir)

	macBookPro13, err := factory.CreateNewMac(factory.MacBookPro13)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MacBookPro13:\n%+v\n", macBookPro13)

	macBookPro16, err := factory.CreateNewMac(factory.MacBookPro16)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MacBookPro16:\n%+v\n", macBookPro16)
}
