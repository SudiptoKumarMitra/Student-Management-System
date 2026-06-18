package main

import (
	student "6th/Student"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("-------Welcome To Student Management System---------")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Student")
		fmt.Println("3. Search Student Via Id")
		fmt.Println("4. Update Information")
		fmt.Println("5. Delete Data")
		fmt.Println("6. Save Data")
		fmt.Println("7. Load Data")
		fmt.Println("8. Exit")
		choicestr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error")
			continue
		}
		choicestr = strings.TrimSpace(choicestr)
		choice, err := strconv.Atoi(choicestr)
		if err != nil {
			fmt.Println("Invalid choice; please enter a number between 1-8")
			continue
		}
		if choice == 1 {
			fmt.Println("Enter Student Name: ")
			sName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error")
				continue
			}
			sName = strings.TrimSpace(sName)
			fmt.Println("Enter Age: ")
			sAgest, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error")
				continue
			}
			sAgest = strings.TrimSpace(sAgest)
			sAge, err := strconv.Atoi(sAgest)
			if err != nil {
				fmt.Println("Invalid age; please enter a valid number")
				continue
			}
			fmt.Println("Enter Your Id: ")
			sId, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error")
				continue
			}
			sId = strings.TrimSpace(sId)
			Stu1 := student.Student{Name: sName, Age: sAge, Id: sId}
			if err := student.AddStudent(Stu1); err != nil {
				fmt.Println("Add failed:", err)
			} else {
				fmt.Println("Added Successfully!")
			}
		} else if choice == 2 {
			student.ShowStudentDetails()
		} else if choice == 3 {
			fmt.Println("Enter Id to search: ")
			SearchID, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error")
				continue
			}
			SearchID = strings.TrimSpace(SearchID)
			if SearchID == "" {
				fmt.Println("Id cannot be empty!")
				continue
			}
			point, err := student.SearchStudent(SearchID)
			if err != nil {
				fmt.Println("Search failed: ", err)
				continue
			}
			fmt.Println("\n\nName: ", point.Name)
			fmt.Println("Age: ", point.Age)
			fmt.Println("Id: ", point.Id)
			fmt.Println()

		} else if choice == 4 {

			fmt.Println("Enter Your Id: ")
			Id, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error")
				continue
			}
			Id = strings.TrimSpace(Id)
			if Id == "" {
				fmt.Println("Id cannot be empty")
				continue
			}
			if !student.IdExists(Id) {
				fmt.Println("Id not Found!")
				continue
			}
			fmt.Println("Enter New Name: ")
			newName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input Error")
				continue
			}
			if newName == "" {
				fmt.Println("Name cannot be empty")
				continue
			}
			newName = strings.TrimSpace(newName)
			fmt.Println("Enter New Age: ")
			newAgest, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input Error")
				continue
			}
			newAgest = strings.TrimSpace(newAgest)
			newAge, err := strconv.Atoi(newAgest)
			if err != nil {
				fmt.Println("Invalid age; please enter a valid number")
				continue
			}
			if newAge <= 0 {
				fmt.Println("Age must be a positive integer")
				continue
			}
			if err := student.UpdateStudent(Id, newName, newAge); err != nil {
				fmt.Println("Update failed:", err)
				continue
			}
			fmt.Println("Successfully Updated!")

		} else if choice == 5 {
			fmt.Println("Enter Id to delete: ")
			deleteID, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input Error")
				continue
			}
			deleteID = strings.TrimSpace(deleteID)
			if deleteID == "" {
				fmt.Println("Id cannot be empty")
				continue
			}

			fmt.Println("Confirm delete? (1=Yes, 2=No)")
			confStr, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input Error")
				continue
			}
			confirm, err := strconv.Atoi(strings.TrimSpace(confStr))
			if err != nil {
				fmt.Println("Invalid choice; enter 1 or 2")
				continue
			}
			if err := student.DeleteData(deleteID, confirm == 1); err != nil {
				fmt.Println("Delete failed:", err)
			} else {
				fmt.Println("Deleted successfully!")
			}
		} else if choice == 6 {
			if err := student.SaveData(); err != nil {
				fmt.Println("Save failed:", err)
			} else {
				fmt.Println("Saved Successfully!")
			}
		} else if choice == 7 {
			if err := student.LoadData(); err != nil {
				fmt.Println("Load failed:", err)
			}
		} else if choice == 8 {
			fmt.Println("Thanks For using our Project")
			return
		} else {
			fmt.Println("Choose number between 1-8")
		}
	}
}
