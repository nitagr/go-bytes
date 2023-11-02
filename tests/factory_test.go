package tests

import (
	"geektrust/types"
	"testing"
)

func TestAdd(t *testing.T) {
	var courses []types.Course
	courseEmployeeRegMap := make(map[string]types.CourseData)
	courseRegIdMap := make(map[string]string)

	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
