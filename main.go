package main

import (
	"fmt"
	"github.com/larjudge/protocol-buffer-tutorial/pkg/addressbook"
)

func main() {
	p := addressbook.Person{
		Id:    1,
		Name:  "Lar Judge",
		Email: "hello@lar.dev",
		Phones: []*addressbook.Person_PhoneNumber{
			{Number: "1234", Type: addressbook.Person_MOBILE},
		},
	}

	fmt.Printf("Person is %+v", p)
}
