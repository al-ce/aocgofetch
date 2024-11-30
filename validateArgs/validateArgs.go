package validateArgs

import (
	"fmt"
	"strconv"
	"time"
)

func getMaxAocYear() int64 {
	now := time.Now()
	year := now.Year()
	if now.Month() == time.December && now.Day() < 26 {
		year--
	}
	return int64(year)
}

const (
	minAocYear = 2015
)

type ArgsState int

const (
	ValidArgs ArgsState = iota
	YearNotAnInt
	DayNotAnInt
	YearInvalid
	DayInvalid
)

var ArgsErrType = map[ArgsState]string{
	YearNotAnInt: "Could not convert year argument to int",
	DayNotAnInt:  "Could not convert day argument to int",
	YearInvalid:  fmt.Sprintf("Year must be between 2015 and %d", getMaxAocYear()),
	DayInvalid:   "Day must be between 1 and 25",
}

func GetYearAndDay(args []string) (int64, int64, ArgsState) {

	yearArg, dayArg := args[0], args[1]

	// Validate year
	year, err := strconv.ParseInt(yearArg, 10, 64)
	if err != nil {
		return -1, -1, YearNotAnInt
	} else if year < minAocYear || year > getMaxAocYear() {
		return -1, -1, YearInvalid
	}

	// Validate day
	day, err := strconv.ParseInt(dayArg, 10, 64)
	if err != nil {
		return -1, -1, DayNotAnInt
	} else if day < 1 || day > 25 {
		return -1, -1, DayInvalid
	}

	return year, day, ValidArgs
}
