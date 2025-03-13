package main

import (
	"fmt"
	"strings"
)

type Person struct {firstName string; lastName string}

func greet(name string, country string) string {
	return fmt.Sprintf("Greetings %v, from %v", name, country)
}

func runUntilEnd() {
	fmt.Println("Cleaning function")
}

func concatStrings(vals ...strings) (string, int) {
	var cat strings.Builder
	defer runUntilEnd()
	for _, val := range vals {
		fmt.Fprint(&cat, val)
	}
	return cat.String()
}

func (p Person) String() string {
	return fmt.Sprintf("First: %v\nLast: %v\n", p.firstName, p.lastName)
}