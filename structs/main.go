package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{email: "abc@gmail.com", zipCode: 202001}}
	fmt.Println(alex)
	//abc:=[]int{}   //Slice Declaration
	//var ab [3]int  //Array declaration
	var saarthak person
	saarthak.firstName = "Saarthak"
	saarthak.lastName = "Gupta"
	fmt.Println(saarthak)
	fmt.Printf("%+v", saarthak)
	alex.updateName("Alexxy")
	//Here Go can automatically recognize alex and its pointer.
	/*
		All the primitive data types and structs require pointer for updation in other functions.
		Slices, maps, channels don't require pointers. Since, they are resizable. A slice basically contains
		pointer to head, length & capacity. So, when it is passed to a function, a copy is made, but this copy refers to same
		pointer to head.
	*/
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
