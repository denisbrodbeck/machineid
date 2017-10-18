package machineid_test

import (
	"fmt"
	"log"

	"github.com/shahaya/machineid"
)

func Example() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
