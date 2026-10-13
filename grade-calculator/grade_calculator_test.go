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
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 41, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeLetterBoundaries(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{90, "A"},
		{89, "B"},
		{80, "B"},
		{79, "C"},
		{70, "C"},
		{69, "D"},
		{60, "D"},
		{59, "F"},
		{0, "F"},
	}

	for _, tc := range tests {
		gradeCalculator := NewGradeCalculator()

		gradeCalculator.AddGrade("assignment", tc.score, Assignment)
		gradeCalculator.AddGrade("exam", tc.score, Exam)
		gradeCalculator.AddGrade("essay", tc.score, Essay)

		actual_value := gradeCalculator.GetFinalGrade()

		if tc.expected != actual_value {
			t.Errorf("Score %d: expected '%s'; got '%s' instead", tc.score, tc.expected, actual_value)
		}
	}
}

func TestGetGradeUsesWeights(t *testing.T) {
	// 100*.5 + 100*.35 + 0*.15 = 85, which is a B.
	// If the essay weight were applied to the wrong category the letter would change.
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam", 100, Exam)
	gradeCalculator.AddGrade("essay", 0, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "B" {
		t.Errorf("Expected 'B'; got '%s' instead", actual_value)
	}
}

func TestGetGradeAveragesMultipleGrades(t *testing.T) {
	// assignment avg 90, exam avg 90, essay avg 90 -> 90 -> A
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment 1", 80, Assignment)
	gradeCalculator.AddGrade("assignment 2", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 85, Exam)
	gradeCalculator.AddGrade("exam 2", 95, Exam)
	gradeCalculator.AddGrade("essay 1", 90, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "A" {
		t.Errorf("Expected 'A'; got '%s' instead", actual_value)
	}
}

func TestGetGradeNoGradesDoesNotPanic(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "F" {
		t.Errorf("Expected 'F' for a calculator with no grades; got '%s' instead", actual_value)
	}
}

func TestAddGradeIgnoresUnknownType(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("mystery", 100, GradeType(99))

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "F" {
		t.Errorf("Expected unknown grade type to be ignored and return 'F'; got '%s' instead", actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, tc := range tests {
		if actual_value := tc.gradeType.String(); actual_value != tc.expected {
			t.Errorf("Expected '%s'; got '%s' instead", tc.expected, actual_value)
		}
	}
}

// Part 7: Pass/Fail grading mode. Passing requires a C or better (70+).

func TestPassFailPass(t *testing.T) {
	gradeCalculator := NewGradeCalculator(PassFail)

	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 85, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "Pass" {
		t.Errorf("Expected 'Pass'; got '%s' instead", actual_value)
	}
}

func TestPassFailFail(t *testing.T) {
	gradeCalculator := NewGradeCalculator(PassFail)

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 41, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "Fail" {
		t.Errorf("Expected 'Fail'; got '%s' instead", actual_value)
	}
}

func TestPassFailBoundaries(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{70, "Pass"},
		{69, "Fail"},
		{100, "Pass"},
		{0, "Fail"},
	}

	for _, tc := range tests {
		gradeCalculator := NewGradeCalculator(PassFail)

		gradeCalculator.AddGrade("assignment", tc.score, Assignment)
		gradeCalculator.AddGrade("exam", tc.score, Exam)
		gradeCalculator.AddGrade("essay", tc.score, Essay)

		if actual_value := gradeCalculator.GetFinalGrade(); actual_value != tc.expected {
			t.Errorf("Score %d: expected '%s'; got '%s' instead", tc.score, tc.expected, actual_value)
		}
	}
}

func TestPassFailNoGradesFails(t *testing.T) {
	gradeCalculator := NewGradeCalculator(PassFail)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "Fail" {
		t.Errorf("Expected 'Fail'; got '%s' instead", actual_value)
	}
}

func TestLetterGradeModeIsDefaultAndExplicit(t *testing.T) {
	defaultCalculator := NewGradeCalculator()
	explicitCalculator := NewGradeCalculator(LetterGrade)

	for _, gc := range []*GradeCalculator{defaultCalculator, explicitCalculator} {
		gc.AddGrade("assignment", 95, Assignment)
		gc.AddGrade("exam", 95, Exam)
		gc.AddGrade("essay", 95, Essay)

		if actual_value := gc.GetFinalGrade(); actual_value != "A" {
			t.Errorf("Expected 'A'; got '%s' instead", actual_value)
		}
	}
}
