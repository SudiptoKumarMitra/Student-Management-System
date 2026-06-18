# Student Management System

This is a beginner Go CLI project for managing student records.
It supports adding, viewing, searching, updating, deleting, saving, and loading students.

## Features

- Add a student with name, age, and id
- View all stored students
- Search a student by id
- Update student name and age by id
- Delete a student by id with confirmation
- Save student data to `Students.json`
- Load student data from `Students.json`

## Run

From the project root:

```bash
go run .
```

## Save / Load

- Choose `Save Data` from the menu to write the current students to `Students.json`
- Choose `Load Data` from the menu to read students from `Students.json`

## Testing

Run package tests with:

```bash
go test ./Student
```

## Notes

- The project uses a simple package-based design:
  - `main.go` handles CLI interaction
  - `Student/` contains student logic and data persistence
- The current implementation is meant for learning Go fundamentals.
