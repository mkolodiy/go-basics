package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) printName() {
	fmt.Println("Name:", p.firstName, p.lastName)
}

func (p person) printContact() {
	fmt.Println("Email:", p.contactInfo.email, "Zip code:", p.contactInfo.zipCode)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	alex := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "john@doe.com",
			zipCode: 1234,
		},
	}
	alex.printName()

	alex.updateName("Alex")
	alex.printName()
}
