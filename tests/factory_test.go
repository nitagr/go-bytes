package tests

import (
	"fmt"
	"geektrust/course"
	"geektrust/errortypes"
	"geektrust/types"
	"testing"
)

func TestAddCourse(t *testing.T) {
	var courses []types.Course
	courseEmployeeRegMap := make(map[string]types.CourseData)
	courseRegIdMap := make(map[string]string)

	parameters := []string{"JAVA", "JAMES", "15062022", "1", "2"}
	result, _ := course.AddCourseOffering(parameters, &courses, courseEmployeeRegMap, courseRegIdMap)

	expected := "OFFERING-JAVA-JAMES"

	if result != expected {
		t.Errorf(" %s; %s", result, expected)
	}
}

func TestEmployeeReg(t *testing.T) {
	var courses []types.Course
	courseEmployeeRegMap := make(map[string]types.CourseData)
	courseRegIdMap := make(map[string]string)

	parameters := []string{"JAVA", "JAMES", "15062022", "1", "2"}
	result, _ := course.AddCourseOffering(parameters, &courses, courseEmployeeRegMap, courseRegIdMap)

	regParams := []string{"nitxyz@gmail.com", result}
	id, status, err := course.RegisterCourse(regParams, &courses, courseEmployeeRegMap, courseRegIdMap)

	expectedregId := "REG-COURSE-nitxyz-JAVA"
	expectedStatus := "ACCEPTED"
	if id != expectedregId && status != expectedStatus {
		fmt.Println(err)
		t.Errorf(" %s %s %s %s", id, expectedregId, status, expectedStatus)
	}
}

func TestCancelReg(t *testing.T) {
	var courses []types.Course
	courseEmployeeRegMap := make(map[string]types.CourseData)
	courseRegIdMap := make(map[string]string)

	regParams := []string{"nitxyz@gmail.com", "OFFERING-JAVA-JAMES"}
	_, _, err := course.RegisterCourse(regParams, &courses, courseEmployeeRegMap, courseRegIdMap)

	if err != errortypes.ErrCourseNotFoundError {
		t.Errorf(" %s %s", err, errortypes.ErrCourseNotFoundError)
	}
}
