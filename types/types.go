package types

type Course struct {
	Id         string
	Name       string
	Instructor string
	Date       string
	MinimumCap int32
	MaximumCap int32
}

type CourseEmployeeRegistrationData struct {
	CourseRegId string
	EmailId     string
	CourseOffId string
	CourseName  string
	Instructor  string
	Date        string
	Status      string
}

type CourseData struct {
	Course
	RegisteredEmployees []CourseEmployeeRegistrationData
	IsAlloted           bool
	IsCanceled          bool
}
