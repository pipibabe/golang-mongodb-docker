package hello

import (
	"fmt"
	"log"
	"rsc.io/quote"
	"golang_test/cmd/hello/sentence"
	"golang_test/internal/greetings"
	// "golang_test/internal/data"
)
func Hello() {
	fmt.Println("Hello, Pipi!")
	// test external package
	fmt.Println(quote.Go())

	// test random sentence
	fmt.Println(sentence.Random())

	// test greetings someone..
	message, err := greetings.Hello("Pipi")
	fmt.Println(message)


	// test empty name
	log.SetPrefix("greetings: ")
	log.SetFlags(83) // log.Ldate(1)|log.Ltime(2)|log.Lshortfile(16)Lmsgprefix(64)
	message, err = greetings.Hello("")
	if err != nil {
		log.Println(err)
		// log.Panic(err)
	}
	log.Println(message)

	// test mongo database
	// data.Insert()


	
}