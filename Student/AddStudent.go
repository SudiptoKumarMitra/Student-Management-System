package student

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
	Id   string
}

var Students []Student

func IdExists(id string) bool {
	for _, i := range Students {
		if i.Id == id {
			return true
		}
	}
	return false
}

func AddStudent(s Student) error {
	if s.Name == "" {
		return fmt.Errorf("Name cannot be empty")
	}
	if s.Age <= 0 {
		return fmt.Errorf("Age must be a positive integer")
	}
	if s.Id == "" {
		return fmt.Errorf("Id cannot be empty")
	}
	if IdExists(s.Id) {
		return fmt.Errorf("Id already exists")
	}

	Students = append(Students, s)
	return nil
}
