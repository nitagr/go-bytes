package course

import (
	"fmt"
	"geektrust/constants"
	"geektrust/errortypes"
	"geektrust/types"
)

func ExecuteCommandsFactory(
	commandText []string,
	courses *[]types.Course,
	courseEmployeeRegMap map[string]types.CourseData,
	courseRegIdMap map[string]string,
) {
	defer (func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	})()
	if len(commandText) <= 0 {
		return
	}

	currentCommand, parameters := commandText[0], commandText[1:]
	errValidCommand := inputCommandValidation(currentCommand, parameters)

	if errValidCommand != nil {
		panic(errValidCommand)
	}

	switch currentCommand {
	case constants.ADD_COURSE_OFFERING:
		addCourseOffering(parameters, courses, courseEmployeeRegMap, courseRegIdMap)

	case constants.REGISTER:
		registerCourse(parameters, courses, courseEmployeeRegMap, courseRegIdMap)

	case constants.ALLOT_COURSE:
		allotCourse(parameters, courses, courseEmployeeRegMap, courseRegIdMap)

	case constants.CANCEL:
		cancelRegistration(parameters, courses, courseEmployeeRegMap, courseRegIdMap)

	default:
		panic(errortypes.ErrCommandNotFoundError)
	}

}