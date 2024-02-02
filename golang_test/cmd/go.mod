module main

go 1.21.6

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace golang_test/cmd/hello => ./hello

replace golang_test/cmd/hello/sentence => ./sentence

replace golang_test/internal/greetings => ../../internal/greetings
