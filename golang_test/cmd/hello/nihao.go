package hello_nihao

import (
	"fmt"
	"log"
	"rsc.io/quote"
	"golang_test/cmd/hello/sentence"
	"golang_test/internal/greetings"
	"golang_test/internal/data"
)
func Hello() {
	fmt.Println("Nihao, Pipi!")
	// test external package
	fmt.Println(quote.Go())

	// test random sentence
	fmt.Println(sentence.Random())

	// test greetings someone..
	message, err := greetings.Hello("Pipi")
	fmt.Println(message)


	// test empty name
	log.SetPrefix("greetings: ")
    log.SetFlags(0)
	message, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// test mongo database
	data.person.GetAdminDatabase()


	
}