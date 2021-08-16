package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
	Person
}

type FullTimeEmployee struct {
	Employee
	Hours int
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployee = func(id int, dni string) (Employee, error) {
	time.Sleep(5 * time.Second)

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return Employee{}, err
	}

	return Employee{
		Person: p,
	}, nil
}

func GetFullTimeEmployee(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployee(id, dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = e

	return ftEmployee, nil
}
