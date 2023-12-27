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
}