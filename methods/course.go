package methods

import (
	"geektrust/constants"
	course "geektrust/errors_types"
)

func InputCommandValidation(currentCommand string, parameters []string) error {
	var paramsLength int
	switch currentCommand {
	case constants.ADD_COURSE_OFFERING:
		paramsLength = len(parameters)
		if paramsLength < 5 {
			return course.ErrCourseFullError
		}

	case constants.ALLOT_COURSE:
		paramsLength = len(parameters)
		if paramsLength < 1 {
			return course.ErrCourseFullError
		}

	case constants.REGISTER:
		paramsLength = len(parameters)
		if paramsLength < 2 {
			return course.ErrCourseFullError
		}

	case constants.CANCEL:
		paramsLength = len(parameters)
		if paramsLength < 1 {
			return course.ErrCourseFullError
		}
	}
	return nil

}
