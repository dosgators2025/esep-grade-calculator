
# ESEP Grade Calculator

this project is a Go program that calculates a student's final letter grade using assignments, exams, and essays.

## Requirements

- Go 1.24.4 or newer
- Git

## Getting Started

clone the repository:


```bash
git clone https://github.com/dosgators2025/esep-grade-calculator.git
```

open the project folder:

```bash
cd esep-grade-calculator/grade-calculator
```

## Grade Calculation

the final grade is calculated using these weights:

- Assignments: 50%
- Exams: 35%
- Essays: 15%

the calculator averages the grades in each category, applies the weights, and returns a letter grade (A, B, C, D, or F).

## Running Tests

run the unit tests from the `grade-calculator` folder:

```bash
go test -v
```

The `-v` option shows the results of each test.

## Code Coverage

to generate a coverage report:

```bash
go test -coverprofile=coverage.txt ./...
```

to view the coverage percentage:

```bash
go tool cover -func=coverage.txt
```
