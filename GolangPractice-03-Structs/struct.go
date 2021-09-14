package main

import "fmt"

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

func main() {

	// You can name the fields when initializing a struct.
	p1 := person{
		name: "Tim",
		age:  35}
	fmt.Println(p1)

	// use shorter syntax to initialize a struct
	// structName{field 1, field 2}
	p2 := person{"John", 12}
	fmt.Println(p2)

	p3 := person{"Alice", 18}
	// Access struct fields with a dot.
	fmt.Println(p3.age)

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Angela"})

	// Structs are mutable.
	p4 := person{"Angela", 20}
	p4.age = 23
	fmt.Println(p4)

	// using the current struct as a field in another struct
	type employee struct {
		person person
		id     int
		salary int
	}

	//populate fields in the new struct
	emp1 := employee{
		person: person{"John", 34},
		id:     1038,
		salary: 3500}

	fmt.Println(emp1)
	// to access fields of nested structs use the following syntax
	fmt.Println(emp1.person.name)

}
