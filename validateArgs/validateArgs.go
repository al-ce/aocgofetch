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

func safeIndex(args []string, index int) string {
	if index >= 0 && index < len(args) {
		return args[index]
	}
	return ""
}

func (self ArgsState) Error(args []string) error {
	year, day := safeIndex(args, 0), safeIndex(args, 1)

	ArgsErrType := map[ArgsState]string{
		BadArgsAmount:  fmt.Sprintf("got %d args, need exactly 2: <day> <year>", len(args)),
		YearArgNotInt:  fmt.Sprintf("%s: could not convert year argument to int", year),
		DayArgNotInt:   fmt.Sprintf("%s: could not convert day argument to int", day),
		DayArgInvalid:  fmt.Sprintf("%s: day must be between 1 and %d", day, getMaxAocDay()),
		YearArgInvalid: fmt.Sprintf("%s: year must be between 2015 and %d", year, getMaxAocYear()),
	}
	if err, exists := ArgsErrType[self]; exists {
		return fmt.Errorf(err)
	}
	return fmt.Errorf("unknown args error")
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
