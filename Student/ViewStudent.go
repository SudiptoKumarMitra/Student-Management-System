package student

import (
	"fmt"
)

func ShowStudentDetails() {
	if len(Students) == 0 {
		fmt.Println("\n\nEnter Atleast One Student Information")
	} else {
		for j, i := range Students {
			fmt.Println(j+1, "No. Student")
			fmt.Println("Name: ", i.Name)
			fmt.Println("Age: ", i.Age)
			fmt.Println("Id: ", i.Id)
			fmt.Println()
			fmt.Println()
		}
	}
}
