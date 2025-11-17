package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day byte
	Month byte
	Year int
}

type Profile struct {
	// TODO: embed full name and birth date information
	NumberOfSiblings byte
	ZodiacSign       rune
	Name FullName
	BirthDate BirthDate
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		NumberOfSiblings: 3,   // TODO: adjust
		ZodiacSign:       '\u2649', // TODO: adjust
		Name: FullName{
			FirstName: "Rafael",
			LastName: "Hertling",
		},
		BirthDate: BirthDate{
			Day: 02,
			Month: 05,
			Year: 2009,
		},
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	fmt.Println("Siblings After:", me.NumberOfSiblings+1)
} 