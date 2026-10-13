package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam", 75, Exam)
	gradeCalculator.AddGrade("essay", 75, Essay)

	expected := "C"
	actual := gradeCalculator.GetFinalGrade()

	if expected != actual {
		t.Errorf("Expected C, got %s", actual)
	}
}

func TestGradeTypeString(t *testing.T) {
	expected := "assignment"
	actual := Assignment.String()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGetGradeD(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam", 65, Exam)
	gradeCalculator.AddGrade("essay", 65, Essay)

	expected := "D"
	actual := gradeCalculator.GetFinalGrade()

	if expected != actual {
		t.Errorf("Expected D, got %s", actual)
	}
}
