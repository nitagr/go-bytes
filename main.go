package main

import (
	"bufio"
	"fmt"
	"geektrust/constants"
	"geektrust/types"
	"geektrust/util"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var courses []types.Course
var courseRegistrationMap map[string]types.CourseStatusData
var courseRegIdMap map[string]string

func sortByRegistrationId(registrations []types.CourseRegistrationData) {
	sort.Slice(registrations, func(i, j int) bool {
		return registrations[i].CourseRegId < registrations[j].CourseRegId
	})
}

func removeRegistrationById(registrations []types.CourseRegistrationData, courseRegId string) []types.CourseRegistrationData {
	var indexToRemove int
	for i, reg := range registrations {
		if reg.CourseRegId == courseRegId {
			indexToRemove = i
			break
		}
	}

	// Remove the element by slicing the slice
	return append(registrations[:indexToRemove], registrations[indexToRemove+1:]...)
}

func inputCommandValidation(currentCommand string, parameters []string) bool {
	var paramsLength int
	switch currentCommand {
	case constants.ADD_COURSE_OFFERING:
		paramsLength = len(parameters)
		if paramsLength < 5 {
			return false
		}

	case constants.ALLOT_COURSE:
		paramsLength = len(parameters)
		if paramsLength < 1 {
			return false
		}

	case constants.REGISTER:
		paramsLength = len(parameters)
		if paramsLength < 2 {
			return false
		}

	case constants.CANCEL:
		paramsLength = len(parameters)
		if paramsLength < 1 {
			return false
		}
	}
	return true

}

func executeCommands(
	commandText []string,
	// courses *[]types.Course,
	// courseEnrollmentMap map[string]types.CourseStatusMetaData,
) {
	currentCommand, parameters := commandText[0], commandText[1:]
	isValidCommand := inputCommandValidation(currentCommand, parameters)

	if !isValidCommand {
		fmt.Println(constants.INPUT_DATA_ERROR)
		return
	}
	switch currentCommand {
	case constants.ADD_COURSE_OFFERING:

		name, instructor, courseDate, minCapacity, maxCapacity := parameters[0], parameters[1], parameters[3], parameters[3], parameters[4]
		minCap, _ := strconv.Atoi(minCapacity)
		maxCap, _ := strconv.Atoi(maxCapacity)

		courseOfferingId := constants.OFFERING + "-" + name + "-" + instructor
		formattedDate := util.DateFormat(courseDate)

		course := types.Course{
			Id:         courseOfferingId,
			Name:       name,
			Instructor: instructor,
			Date:       formattedDate,
			MinimumCap: int32(minCap),
			MaximumCap: int32(maxCap),
		}

		if courseRegistrationMap == nil {
			courseRegistrationMap = make(map[string]types.CourseStatusData)
		}

		if _, ok := courseRegistrationMap[courseOfferingId]; !ok {
			courseRegistrationMap[courseOfferingId] = types.CourseStatusData{
				Course:              course,
				RegisteredEmployees: []types.CourseRegistrationData{},
				IsAlloted:           false,
				IsCanceled:          false,
			}

			courses = append(courses, course)

		}

	case constants.REGISTER:
		email, courseOfferingId := parameters[0], parameters[1]
		employeeName := strings.Split(email, "@")[0]
		courseNameIns := strings.Split(courseOfferingId, "-")
		courseName, instructor := courseNameIns[1], courseNameIns[2]

		if course, ok := courseRegistrationMap[courseOfferingId]; ok {
			if ok {
				registeredEmployees := course.RegisteredEmployees
				maxCap := course.Course.MaximumCap

				if len(registeredEmployees) < int(maxCap) {
					registrationId := constants.REG_COURSE + "-" + employeeName + "-" + courseName
					employeeRegData := types.CourseRegistrationData{
						CourseRegId: registrationId,
						EmailId:     email,
						CourseOffId: courseOfferingId,
						CourseName:  courseName,
						Instructor:  instructor,
						Date:        time.Now(),
						Status:      constants.ACCEPTED,
					}

					if courseRegIdMap == nil {
						courseRegIdMap = make(map[string]string)
					}

					courseRegValue := courseRegistrationMap[courseOfferingId].RegisteredEmployees
					courseRegValue = append(courseRegValue, employeeRegData)
					course.RegisteredEmployees = courseRegValue
					courseRegistrationMap[courseOfferingId] = course

					courseRegIdMap[registrationId] = courseOfferingId
					fmt.Println(registrationId, " ", constants.ACCEPTED)

				} else {
					fmt.Println(" ", constants.COURSE_FULL_ERROR)
				}
			}

		}

	case constants.ALLOT_COURSE:
		courseOfferingId := parameters[0]
		courseData := courseRegistrationMap[courseOfferingId]

		registrations := courseData.RegisteredEmployees
		if len(registrations) < int(courseData.Course.MinimumCap) {
			fmt.Println(constants.COURSE_CANCELED)
			courseData.IsCanceled = true
			courseRegistrationMap[courseOfferingId] = courseData
			return
		}
		sortByRegistrationId(registrations)
		courseData.IsAlloted = true
		fmt.Printf("%+v\n", registrations)
		courseRegistrationMap[courseOfferingId] = courseData

	case constants.CANCEL:
		courseRegId := parameters[0]
		courseOfferingId := courseRegIdMap[courseRegId]

		courseData := courseRegistrationMap[courseOfferingId]
		if courseData.IsAlloted {
			fmt.Println(courseRegId, " ", constants.CANCEL_REJECTED)
		} else {
			updatedEmployees := removeRegistrationById(courseData.RegisteredEmployees, courseRegId)
			courseData.RegisteredEmployees = updatedEmployees
			courseRegistrationMap[courseOfferingId] = courseData

			fmt.Println(courseRegId, " ", constants.CANCEL_ACCEPTED)
		}

	}

}

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		args := scanner.Text()
		argList := strings.Fields(args)
		executeCommands(argList)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
