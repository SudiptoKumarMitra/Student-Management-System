package student

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadData() error {
	data, err := os.ReadFile("Students.json")
	if err != nil {
		return fmt.Errorf("read file failed: %w", err)
	}

	if err := json.Unmarshal(data, &Students); err != nil {
		return fmt.Errorf("json unmarshal failed: %w", err)
	}
	ShowStudentDetails()
	fmt.Println("Loaded Successfully!")
	return nil
}
