package course

import (
	"fmt"
	"geektrust/constants"
	"geektrust/errortypes"
	"geektrust/types"
	"strconv"
	"strings"
)

func addCourseOffering(
	parameters []string,
	courses *[]types.Course,
	courseEmployeeRegMap map[string]types.CourseData,
	courseRegIdMap map[string]string,
) {
	name, instructor, courseDate, minCapacity, maxCapacity := parameters[0], parameters[1], parameters[2], parameters[3], parameters[4]
	minCap, _ := strconv.Atoi(minCapacity)
	maxCap, _ := strconv.Atoi(maxCapacity)

	courseOfferingId := constants.OFFERING + "-" + name + "-" + instructor
	course := types.Course{
		Id:         courseOfferingId,
		Name:       name,
		Instructor: instructor,
		Date:       courseDate,
		MinimumCap: int32(minCap),
		MaximumCap: int32(maxCap),
	}

	if _, ok := courseEmployeeRegMap[courseOfferingId]; !ok {
		courseEmployeeRegMap[courseOfferingId] = types.CourseData{
			Course:              course,
			RegisteredEmployees: []types.CourseEmployeeRegistrationData{},
			IsAlloted:           false,
			IsCanceled:          false,
		}

		*courses = append(*courses, course)
		fmt.Println(courseOfferingId)

	} else {
		panic(errortypes.ErrCourseAlreadyExistsError)
	}

}

func registerCourse(
	parameters []string,
	courses *[]types.Course,
	courseEmployeeRegMap map[string]types.CourseData,
	courseRegIdMap map[string]string,
) {
	email, courseOfferingId := parameters[0], parameters[1]
	employeeName := strings.Split(email, "@")[0]
	courseNameIns := strings.Split(courseOfferingId, "-")
	courseName, instructor := courseNameIns[1], courseNameIns[2]

	if course, ok := courseEmployeeRegMap[courseOfferingId]; ok {

		registeredEmployees := course.RegisteredEmployees
		maxCap := course.Course.MaximumCap

		if len(registeredEmployees) < int(maxCap) {
			registrationId := constants.REG_COURSE + "-" + employeeName + "-" + courseName
			employeeRegData := types.CourseEmployeeRegistrationData{
				CourseRegId: registrationId,
				EmailId:     email,
				CourseOffId: courseOfferingId,
				CourseName:  courseName,
				Instructor:  instructor,
				Date:        course.Course.Date,
				Status:      constants.CONFIRMED,
			}

			courseRegValue := courseEmployeeRegMap[courseOfferingId].RegisteredEmployees
			courseRegValue = append(courseRegValue, employeeRegData)
			course.RegisteredEmployees = courseRegValue
			courseEmployeeRegMap[courseOfferingId] = course

			courseRegIdMap[registrationId] = courseOfferingId
			fmt.Println(registrationId, constants.ACCEPTED)

		} else {
			fmt.Println(constants.COURSE_FULL_ERROR)
		}

	} else {
		panic(errortypes.ErrCourseNotFoundError)
	}
}

func allotCourse(
	parameters []string,
	courses *[]types.Course,
	courseEmployeeRegMap map[string]types.CourseData,
	courseRegIdMap map[string]string,
) {

	courseOfferingId := parameters[0]
	courseData, ok := courseEmployeeRegMap[courseOfferingId]

	if !ok {
		fmt.Println("")
		return
	}

	registrations := courseData.RegisteredEmployees
	if len(registrations) < int(courseData.Course.MinimumCap) {
		fmt.Println(constants.COURSE_CANCELED)
		courseData.IsCanceled = true
		courseEmployeeRegMap[courseOfferingId] = courseData
		return
	}
	sortByRegistrationId(registrations)
	courseData.IsAlloted = true
	courseEmployeeRegMap[courseOfferingId] = courseData
	listRegisteredEmployees(registrations)

}

func cancelRegistration(
	parameters []string,
	courses *[]types.Course,
	courseEmployeeRegMap map[string]types.CourseData,
	courseRegIdMap map[string]string,
) {

	courseRegId := parameters[0]
	courseOfferingId, ok := courseRegIdMap[courseRegId]
	if !ok {
		fmt.Println("")
		return
	}

	if courseData, ok := courseEmployeeRegMap[courseOfferingId]; ok {
		if courseData.IsAlloted {
			fmt.Println(courseRegId, constants.CANCEL_REJECTED)
		} else {
			updatedEmployees := removeRegistrationById(courseData.RegisteredEmployees, courseRegId)
			courseData.RegisteredEmployees = updatedEmployees
			courseEmployeeRegMap[courseOfferingId] = courseData

			fmt.Println(courseRegId, constants.CANCEL_ACCEPTED)
		}
	}

}
