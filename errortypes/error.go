package errortypes

import "errors"

var ErrCourseFullError = errors.New("COURSE_FULL_ERROR")
var ErrInputDataError = errors.New("INPUT_DATA_ERROR")
var ErrCommandNotFoundError = errors.New("COMMAND_NOT_FOUND_ERROR")
var ErrCourseAlreadyExistsError = errors.New("COURSE_ALREADY_EXISTS_ERROR")
var ErrCourseNotFoundError = errors.New("COURSE_NOT_FOUND_ERROR")
