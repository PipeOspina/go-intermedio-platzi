package main

import (
	"testing"
)

func TestGetFullTimeEmployee(t *testing.T) {
	table := []struct {
		id       int
		dni      string
		mockFunc func()
		expected FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployee = func(id int, dni string) (Employee, error) {
					p, err := GetPersonByDNI(dni)

					if err != nil {
						return Employee{}, err
					}

					return Employee{
						Id:       1,
						Position: "CEO",
						Person:   p,
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Pipe",
						Age:  23,
						DNI:  "1",
					}, nil
				}
			},
			expected: FullTimeEmployee{
				Employee: Employee{
					Id:       1,
					Position: "CEO",
					Person: Person{
						DNI:  "1",
						Name: "Pipe",
						Age:  23,
					},
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployee
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()

		ft, err := GetFullTimeEmployee(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when gettin FullTimeEmployee")
		}

		if ft != test.expected {
			t.Errorf("Error, got %+v expected %+v", ft, test.expected)
		}
	}

	GetEmployee = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
