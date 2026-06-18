package student

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveData() error {
	data, err := json.MarshalIndent(Students, "", "    ")
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	if err := os.WriteFile("Students.json", data, 0644); err != nil {
		return fmt.Errorf("write file failed: %w", err)
	}

	return nil
}
