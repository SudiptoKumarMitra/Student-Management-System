package student

import (
	"fmt"
)

func DeleteData(id string, confirmed bool) error {
	if !confirmed {
		return fmt.Errorf("Deletion cancelled")
	}

	for i := range Students {
		if Students[i].Id == id {
			Students = append(Students[:i], Students[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Student with id '%s' not found", id)
}
