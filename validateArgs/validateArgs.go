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

func (self ArgsState) Error(args []string) error {
	ArgsErrType := map[ArgsState]string{
		BadArgsAmount:  fmt.Sprintf("got %d args, need exactly 2: <day> <year>", len(args)),
		YearArgNotInt:  fmt.Sprintf("%s: could not convert year argument to int", args[0]),
		DayArgNotInt:   fmt.Sprintf("%s: could not convert day argument to int", args[1]),
		DayArgInvalid:  fmt.Sprintf("%s: day must be between 1 and %d", args[1], getMaxAocDay()),
		YearArgInvalid: fmt.Sprintf("%s: year must be between 2015 and %d", args[0], getMaxAocYear()),
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
