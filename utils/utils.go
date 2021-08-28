package utils

import (
	"strconv"
	"time"
)

func (dateHelper *DateHelper) AgeFromDate(time time.Time) (int, error) {

	timeNow := dateHelper.Now()
	ageYear := timeNow.Year() - time.Year()

	// Combine the day and month into an integer
	// Eg :-
	// 21 Feb = 212
	// 12 October = 1210
	dobDayMonth, err := strconv.Atoi(strconv.Itoa(time.Day()) + strconv.Itoa(int(time.Month())))
	if err != nil {
		return 0, err
	}
	nowDayMonth, err := strconv.Atoi(strconv.Itoa(timeNow.Day()) + strconv.Itoa(int(timeNow.Month())))
	if err != nil {
		return 0, err
	}

	// if the day + month is larger than today's day + month
	// then the age is still younger by 1 year
	if dobDayMonth > nowDayMonth {
		ageYear = ageYear - 1
	}
	return ageYear, nil
}

func (dateHelper *DateHelper) Parse(layout, date string) (time.Time, error) {
	return time.Parse(layout, date)
}

func (dateHelper *DateHelper) Now() time.Time {
	return time.Now()
}

func (dateHelper *DateHelper) Hours(sec int) int {
	return sec/60/60
}

func (dateHelper *DateHelper) Days(sec int) int {
	return dateHelper.Hours(sec) / 24
}
