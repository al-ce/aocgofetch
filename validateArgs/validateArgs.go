package validateArgs

import (
	"fmt"
	"strconv"
	"time"
)

const (
	minAocYear = 2015
)

type ArgsState int

const (
	ValidArgs ArgsState = iota
	BadArgsAmount
	YearArgNotInt
	DayArgNotInt
	YearArgInvalid
	DayArgInvalid
)

var ArgsErrType = map[ArgsState]string{
	BadArgsAmount:  "need exactly two args: <day> <year>",
	YearArgNotInt:  "could not convert year argument to int",
	DayArgNotInt:   "could not convert day argument to int",
	DayArgInvalid:  fmt.Sprintf("day must be between 1 and %d", getMaxAocDay()),
	YearArgInvalid: fmt.Sprintf("year must be between 2015 and %d", getMaxAocYear()),
}

func getMaxAocYear() int {
	now := time.Now()
	maxYear := now.Year()
	if now.Month() < time.December {
		maxYear--
	}
	return maxYear
}

func getMaxAocDay() int {
	maxYear, maxDay := getMaxAocYear(), 25
	thisYear, today := time.Now().Year(), time.Now().Day()
	if thisYear == maxYear && today <= 25 {
		maxDay = today
	}
	return maxDay
}

func GetYearAndDay(args []string) (string, string, ArgsState) {
	if len(args) != 2 {
		return "", "", BadArgsAmount
	}

	yearArg, dayArg := args[0], args[1]

	if yearInt, err := strconv.Atoi(yearArg); err != nil {
		return "", "", YearArgNotInt
	} else if yearInt < minAocYear || yearInt > getMaxAocYear() {
		return "", "", YearArgInvalid
	}

	if dayInt, err := strconv.Atoi(dayArg); err != nil {
		return "", "", DayArgNotInt
	} else if dayInt < 1 || dayInt > getMaxAocDay() {
		return "", "", DayArgInvalid
	}

	return yearArg, dayArg, ValidArgs
}
