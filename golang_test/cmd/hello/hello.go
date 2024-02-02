package hello

import (
	"fmt"
	"log"
	"rsc.io/quote"
	"golang_test/cmd/hello/sentence"
	"golang_test/internal/greetings"
)
func Hello() {
	fmt.Println("Hello, Pipi!")
	fmt.Println(quote.Go())
	fmt.Println(sentence.Random())

	//test greetings someone..
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

	
}