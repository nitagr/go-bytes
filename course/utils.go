package course

import (
	"fmt"
	"geektrust/constants"
	errorTypes "geektrust/errortypes"
	"geektrust/types"
	"sort"
)

// sorting in ascending order of registration ids
func sortByRegistrationId(registrations []types.CourseEmployeeRegistrationData) {
	sort.Slice(registrations, func(i, j int) bool {
		return registrations[i].CourseRegId < registrations[j].CourseRegId
	})
}

// remove registration by course registration id
func removeRegistrationById(
	registrations []types.CourseEmployeeRegistrationData,
	courseRegId string,
) []types.CourseEmployeeRegistrationData {

	var indexToRemove int
	for i, reg := range registrations {
		if reg.CourseRegId == courseRegId {
			indexToRemove = i
			break
		}
	}
	return append(registrations[:indexToRemove], registrations[indexToRemove+1:]...)
}

func listRegisteredEmployees(list []types.CourseEmployeeRegistrationData) {
	for _, p := range list {
		fmt.Printf("%s %s %s %s %s %s %s\n", p.CourseRegId, p.EmailId, p.CourseOffId, p.CourseName, p.Instructor, p.Date, p.Status)
	}
}

// input command validation factory
func inputCommandValidation(currentCommand string, parameters []string) error {
	paramsLength := len(parameters)
	switch currentCommand {
	case constants.ADD_COURSE_OFFERING:
		if paramsLength < 5 {
			return errorTypes.ErrInputDataError
		}

	case constants.ALLOT_COURSE:
		if paramsLength < 1 {
			return errorTypes.ErrInputDataError
		}

	case constants.REGISTER:
		if paramsLength < 2 {
			return errorTypes.ErrInputDataError
		}

	case constants.CANCEL:
		if paramsLength < 1 {
			return errorTypes.ErrInputDataError
		}

	}
	return nil

}
