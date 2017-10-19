package machineid_test

import (
	"fmt"
	"log"

	"github.com/denisbrodbeck/machineid"
)

func Example() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func Example_protected() {
	appID := "Corp.SomeApp"
	id, err := machineid.ProtectedID(appID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
