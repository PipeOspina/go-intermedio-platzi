package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	id       int
	vacation bool
}

type FullTimeEmployee struct {
	Employee
	hours int
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Pipe"
	ftEmployee.age = 21
	ftEmployee.id = 1214741495
	fmt.Printf("%v\n", ftEmployee)

	ftEmployee.Employee = Employee{
		id:       123,
		vacation: true,
		Person: Person{
			name: "Luis",
			age:  36,
		},
	}
	fmt.Printf("%v\n", ftEmployee)

	ftEmployee.Person = Person{
		name: "Santy",
		age:  20,
	}
	fmt.Printf("%v\n", ftEmployee)
}
