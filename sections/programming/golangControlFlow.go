package main

import "fmt"

func main() {
	age := 16
	counter := 1
	fruits := []string{"apple", "orange", "bananas"}
	surnames := map[string]string {
		"Osvaldo": "Valdo",
		"Israel": "Isra",
	}

	if (0 <= age) && (age <= 14) {
		fmt.Println("kid")
	} else if (14 < age) && (age <= 25) {
		fmt.Println("teenager")
	} else {
		fmt.Println("adult")
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("For loop: %v\n", i)
	}

	for counter <= 20 {
		counter += 1
		fmt.Printf("While loop: %v\n", counter)
	}

	for {
		counter += 1
		fmt.Printf("While true loop: %v\n", counter)
		if counter % 5 == 0 {
			break
		}
	}

	for index, fruit := range fruits {
		fmt.Printf("Fruit %v in index %v\n", fruit, index)
	}

	for name, surname := range surnames {
		fmt.Printf("%v, aka %v\n", name, surname)
	}

}