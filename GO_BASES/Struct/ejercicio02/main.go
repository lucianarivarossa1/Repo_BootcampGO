package main

import "fmt"

// Person is a struct that represent a person for our example
type Person struct {
	//ID is the unique identifier for a person
	ID int
	//Name is the full name of a person
	Name string
	//DateofBirth is the birthday of a person
	DateofBirth string
}

// Employee is a struct that represent a person who is a employee
type Employee struct {
	Person
	//ID is the unique identifier for a employee
	ID int
	//Position is the position of a employee
	Position string
}

// PrintEmployee
func (e Employee) PrintEmployee() {
	fmt.Println()
}
func main() {
	person := Person{
		ID:          123,
		Name:        "Luciana",
		DateofBirth: "3-3-3",
	}
	employee := Employee{
		Person:   person,
		ID:       222,
		Position: "asasas",
	}
	id_person1 := person.Name
	id_employee := employee.ID
	id_person := employee.Person.ID
	fmt.Println("person: ", id_person)
	fmt.Println("employee: ", id_employee)
	fmt.Println("nombre person: ", id_person1)
	employee.PrintEmployee()
}
