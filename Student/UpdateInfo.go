package student

import (
	"fmt"
)

func UpdateStudent(SearchedId string, newName string, newAge int) error {

	for i := range Students {
		p := &Students[i]
		if p.Id == SearchedId {
			p.Name = newName
			p.Age = newAge
			return nil
		}
	}
	return fmt.Errorf("student with id '%s' not found", SearchedId)

}
