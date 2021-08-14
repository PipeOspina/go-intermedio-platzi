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

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
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

	// GetMessage(ftEmployee.Person)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
