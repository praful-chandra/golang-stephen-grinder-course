package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int16
}

type preson struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p preson) print() {
	fmt.Printf("%+v \n", p)
}

func (pointerToPerson *preson) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {

	var person1 preson

	person1.firstName = "Mark"
	person1.lastName = "john"
	person1.contact = contactInfo{
		address: "1st creoss",
		zipCode: 8922,
	}

	// person1.print()

	// person1.contact.address = "2nd street"

	pointerPerson1 := &person1

	pointerPerson1.updateFirstName("updated Name 1")
	person1.print()
	(&person1).updateFirstName("updated Name 2")
	person1.print()
	person1.updateFirstName("Updated name 3")
	person1.print()

}
