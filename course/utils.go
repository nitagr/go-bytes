package course

import (
	"geektrust/constants"
	errorTypes "geektrust/errortypes"
	"geektrust/types"
	"sort"
	"time"
)

func formatDate(date string) time.Time {
	layout := constants.DATE_LAYOUT
	result, _ := time.Parse(layout, date)
	return result
}

func sortByRegistrationId(registrations []types.CourseEmployeeRegistrationData) {
	sort.Slice(registrations, func(i, j int) bool {
		return registrations[i].CourseRegId < registrations[j].CourseRegId
	})
}

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

func inputCommandValidation(currentCommand string, parameters []string) error {
	paramsLength := len(parameters)
	switch currentCommand {
	case constants.ADD_COURSE_OFFERING:
		if paramsLength < 5 {
			return errorTypes.ErrInvalidCommandError
		}

	case constants.ALLOT_COURSE:
		if paramsLength < 1 {
			return errorTypes.ErrCourseFullError
		}

	case constants.REGISTER:
		if paramsLength < 2 {
			return errorTypes.ErrCourseFullError
		}

	case constants.CANCEL:
		if paramsLength < 1 {
			return errorTypes.ErrCourseFullError
		}

	}
	return nil

}
