package student

import "fmt"

func SearchStudent(id string) (*Student, error) {
	for i := range Students {
		if Students[i].Id == id {
			return &Students[i], nil
		}
	}
	return nil, fmt.Errorf("Student with %s Id not found", id)

}
