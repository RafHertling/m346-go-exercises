package main

import "fmt"

func main() {
	const firstName string = "Rafael"
	const lastName string = "Hertling"
	const age int = 16
	const city string = "Luzern"
	const zodiacSign string = "\u2649"

	fmt.Printf("My name is %s %s.\n", firstName, lastName)
	fmt.Printf("I am %d years old.\n", age)
	fmt.Printf("I live in %s.\n", city)
	fmt.Printf("My zodiac sign is %s.\n", zodiacSign)
}
