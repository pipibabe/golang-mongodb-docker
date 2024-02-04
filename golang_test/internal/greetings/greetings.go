package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
/*
list := []string {
	"Alice",
	"Brian",
	"Cindy",
	"Debby",
	"Ethan",
	"Flora",
	"Gebby",
	"Helen",
	"Iris",
	"Jack",
	"Kenny",
	"Lily",
	"Momo",
	"Ninja",
	"Owen",
}
*/