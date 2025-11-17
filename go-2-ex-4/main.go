package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	modules := make(map[int][]Class)

	// Create sample data
	student1 := Student{"Max", "Mustermann"}
	student2 := Student{"Anna", "Schmidt"}
	class1 := Class{[]Student{student1, student2}}

	modules[114] = []Class{class1}
	modules[254] = []Class{class1}
	modules[346] = []Class{class1}

	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
