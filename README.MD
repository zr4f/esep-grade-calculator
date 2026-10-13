# ESEP Grade Calculator

A small Go library that computes a student's final grade in a course.

A student can earn three kinds of grades: assignments, exams, and essays.
The library averages each kind, multiplies each average by its weight, adds the
results, and converts the number to a letter grade.

## Requirements

- Go 1.24 or newer (see `grade-calculator/go.mod`)
- No third party dependencies

## Project layout

```
README.MD                              This file
grade-calculator/
  go.mod                               Go module definition (esep/grade-calculator)
  grade_calculator.go                  GradeCalculator code
  grade_calculator_test.go             Unit tests
```

Note that the folder is named `grade-calculator` (with a hyphen). The Go package
inside it is named `esepunittests`.

## Basic usage

```go
gc := NewGradeCalculator()
gc.AddGrade("homework 1", 95, Assignment)
gc.AddGrade("midterm", 88, Exam)
gc.AddGrade("essay 1", 90, Essay)
fmt.Println(gc.GetFinalGrade()) // prints a letter such as "A"
```

Grade types are `Assignment`, `Exam`, and `Essay`.

## Run the unit tests

```
cd grade-calculator/
go test ./...
```

Add `-v` to see each test by name:

```
go test -v ./...
```

## Run the unit tests with code coverage

Print a coverage percentage:

```
cd grade-calculator/
go test -cover ./...
```

Write a coverage profile and look at it line by line:

```
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt          # coverage per function
go tool cover -html=coverage.txt          # opens a colored report in your browser
```

## Helpful links

- Go testing package: https://pkg.go.dev/testing
- Go coverage blog post: https://go.dev/blog/cover
