package util

import (
	"geektrust/constants"
	"time"
)

func DateFormat(date string) time.Time {

	layout := constants.DATE_LAYOUT

	result, _ := time.Parse(layout, date)

	return result
}
