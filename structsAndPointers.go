package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson() *person {
	p := person{name: "bob"}
	p.age = 30
	return &p
}

func main() {
	fmt.Println("Hello, 世界")
	p := newPerson()
	p.age = 10
	fmt.Println(p.age)
}
