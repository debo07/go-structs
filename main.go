package main

import "fmt"

type address struct {
	lane string
	city string
	state string
	pinCode int
}

type person struct {
	firstName string
	lastName string
}

type personWithAddress struct {
	person
	address
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
	fmt.Printf("%+v \n", p3)

	// Update fields of a struct
	p3.firstName = "John"
	p3.lastName = "Doe"
	fmt.Printf("%+v \n", p3)

	// Nested struct
	// Notice state is not set, so zero value is assigned
	p4 := personWithAddress {
		person: person{
			firstName: "Debajit",
			lastName: "Majumder",
		},
		address: address{
			lane: "Street 1",
			city: "Pune",
			pinCode: 411080,
		},
	}
	fmt.Println(p4.firstName)
	fmt.Println(p4.person.firstName)
	fmt.Printf("%+v \n", p4)

	p4.print()
	p4.updateName("Debo1")
	p4.print()
	fmt.Println(p4.firstName)
	fmt.Println(p4.person.firstName)
}

// Receiver function on struct
func (p personWithAddress) print() {
	fmt.Printf("%+v \n", p)
}

func (p *personWithAddress) updateName(newName string) {
	//(*p).person.firstName = newName
	(*p).firstName = newName
}