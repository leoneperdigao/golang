package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	alex.setName("teste")
	alex.toString()
}

func (p *person) setName(firstName string) {
	(*p).firstName = firstName
}

func (p person) toString() {
	fmt.Println(p)
}
