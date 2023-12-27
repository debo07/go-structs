package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// 1 way of using struct - relies on order of the property - not a good practice
	p1 := person{"Debajit", "Majumder"}
	fmt.Println(p1)

	// 2nd way of using struct
	p2 := person{firstName: "Debo", lastName: "Majumdar"}
	fmt.Println(p2)

	// 3rd way
	var p3 person
	fmt.Println(p3)
	// %+v prints all the key name e.g. {firstName: lastName: }
	fmt.Printf("%+v", p3)

	// Update fields of a struct
	p3.firstName = "John"
	p3.lastName = "Doe"
	fmt.Printf("%+v", p3)
}