package student

import "testing"

func resetStudents() {
	Students = nil
}

func TestAddStudent(t *testing.T) {
	resetStudents()
	err := AddStudent(Student{Name: "Test", Age: 20, Id: "1"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(Students) != 1 {
		t.Fatalf("expected 1 student, got %d", len(Students))
	}
}

func TestSearchStudent(t *testing.T) {
	resetStudents()
	AddStudent(Student{Name: "Test", Age: 20, Id: "1"})

	student, err := SearchStudent("1")
	if err != nil {
		t.Fatalf("expected student, got error %v", err)
	}
	if student.Name != "Test" {
		t.Fatalf("expected name Test, got %s", student.Name)
	}
}

func TestUpdateStudent(t *testing.T) {
	resetStudents()
	AddStudent(Student{Name: "Test", Age: 20, Id: "1"})

	err := UpdateStudent("1", "NewName", 22)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if Students[0].Name != "NewName" || Students[0].Age != 22 {
		t.Fatalf("expected updated student, got %+v", Students[0])
	}
}

func TestDeleteData(t *testing.T) {
	resetStudents()
	AddStudent(Student{Name: "Test", Age: 20, Id: "1"})

	err := DeleteData("1", true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(Students) != 0 {
		t.Fatalf("expected 0 students, got %d", len(Students))
	}
}
