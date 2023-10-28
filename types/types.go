package types

import "time"

type Course struct {
	Id         string
	Name       string
	Instructor string
	Date       time.Time
	MinimumCap int32
	MaximumCap int32
}
