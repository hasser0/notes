package main

import "fmt"

const PiExported = 3.14159

func main() {
	var trueBoolVar bool = true
	var falseBoolVar bool = true
	var stringVar string = "stringVar"
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64
	var intVar int = 50
	// float32, float64
	var floatVar float = 2.7172
	// complex64, complex128
	var complexVar complex64 = 1 + 1i

	// type inference
	intInferred := 10
	floatInferred := 9.81
	boolInferred := true
	stringInferred := "string"
}