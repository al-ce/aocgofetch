package parser

import (
	"fmt"
	"strconv"
	"time"
)

const (
	minAocYear = 2015
)

type ArgsErr int

const (
	ValidArgs ArgsErr = iota
	BadArgsAmount
	YearArgNotInt
	DayArgNotInt
	YearArgInvalid
	DayArgInvalid
)

type ParsedArgs struct {
	Year    int
	Day     int
	MaxYear int
	MaxDay  int
	arg1    string
	arg2    string
	Length  int
}

func getCurrentTimeEST() time.Time {
	loc, err := time.LoadLocation("America/New_York") // AoC releases 12AM EST
	if err != nil {
		panic("aocgofetch: could not load timezone location")
	}

	return time.Now().In(loc)
}

// Get the latest AoC year
// The max AoC year can't be greater than last year if it is not yet December,
// since it would not have been released. Otherwise it is this year.
func getMaxAocYear() int {
	now := getCurrentTimeEST()
	maxYear := now.Year()
	if now.Month() < time.December {
		maxYear--
	}
	return maxYear
}

// Get the latest possible AoC day.
// The max AoC day can't be greater than 25 or greater than the current
// day (in December) if the most recent AoC year is this year
func getMaxAocDay(year int) int {
	maxDay := 25
	now := getCurrentTimeEST()
	thisYear, today := now.Year(), now.Day()
	if thisYear == year && today <= 25 {
		maxDay = today
	}
	return maxDay
}

func Parse(args []string) (ParsedArgs, ArgsErr) {
	parsed := ParsedArgs{-1, -1, getMaxAocYear(), -1, "", "", len(args)}

	if parsed.Length != 2 {
		return parsed, BadArgsAmount
	}

	parsed.arg1, parsed.arg2 = args[0], args[1]

	yearInt, err := strconv.Atoi(parsed.arg1)
	if err != nil {
		return parsed, YearArgNotInt
	}

	parsed.Year = yearInt

	if yearInt < minAocYear || yearInt > parsed.MaxYear {
		return parsed, YearArgInvalid
	}

	parsed.MaxDay = getMaxAocDay(yearInt)

	dayInt, err := strconv.Atoi(parsed.arg2)
	if err != nil {
		return parsed, DayArgNotInt
	}

	parsed.Day = dayInt

	if dayInt < 1 || dayInt > parsed.MaxDay {
		return parsed, DayArgInvalid
	}

	return parsed, ValidArgs
}

func (p ParsedArgs) FmtArgsErr(ae ArgsErr) error {
	ArgsErrType := map[ArgsErr]string{
		BadArgsAmount:  fmt.Sprintf("got %d args, need exactly 2: <day> <year>", p.Length),
		YearArgNotInt:  fmt.Sprintf("%s: <year> is not an int", p.arg1),
		DayArgNotInt:   fmt.Sprintf("%s: <day> is not an int", p.arg2),
		DayArgInvalid:  fmt.Sprintf("%d: <day> must be between 1 and %d", p.Day, p.MaxDay),
		YearArgInvalid: fmt.Sprintf("%d: <year> must be between %d and %d", minAocYear, p.Year, p.MaxDay),
	}

	if err, exists := ArgsErrType[ae]; exists {
		return fmt.Errorf(err)
	}
	return fmt.Errorf("unknown args error")
}
