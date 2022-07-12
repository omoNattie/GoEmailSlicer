package main

import (
	"errors"
	"fmt"
	"strings"
)

var errNoSymbol = errors.New("No @ symbol given.")

func slice(email string) (string, error) {

	if !strings.Contains(email, "@") {
		return "", errNoSymbol
	}

	splitEmail := strings.Split(email, "@")

	name := splitEmail[0]
	domain := splitEmail[1]

	return "Name: " + name + "\n" + "Domain:" + domain + "\n", errors.New("No error returned")
}

func main() {
	fmt.Println(slice("test@example.com"))
	fmt.Println("------")
	fmt.Println(slice("testexample.com"))
}
