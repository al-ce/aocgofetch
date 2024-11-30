package parser

import (
	"fmt"
	"testing"
	"time"
)

func TestValidArgsErr(t *testing.T) {
	args := []string{"2015", "4"}
	_, actual := Parse(args)
	if actual != ValidArgs {
		testArgsErrHelper(t, args, actual)
	}
}

func TestValidArgsParsed(t *testing.T) {
	expected := ParsedArgs{
		Year:    2015,
		Day:     4,
		MaxYear: getMaxAocYear(),
		MaxDay:  getMaxAocDay(2015),
		arg1:    "2015",
		arg2:    "4",
		Length:  2,
	}

	args := []string{"2015", "4"}
	actual, _ := Parse(args)
	if actual != expected {
		t.Fatalf(`parser.Parse(%q) = %v, _, want %v"`, args, actual, expected)
	}
}

func testArgsErrHelper(t *testing.T, args []string, expected ArgsErr) {
	if _, actual := Parse(args); actual != expected {
		t.Fatalf(`parser.Parse(%q) = _, %v, want %v"`, args, actual, expected)
	}
}

func TestTooManyArgs(t *testing.T) {
	args := []string{"2015", "4", "1225"}
	testArgsErrHelper(t, args, BadArgsAmount)
}

func TestTooFewArgs(t *testing.T) {
	args := []string{"2015"}
	testArgsErrHelper(t, args, BadArgsAmount)
}

func TestYearArgNotInt(t *testing.T) {
	args := []string{"notanint", "4"}
	testArgsErrHelper(t, args, YearArgNotInt)
}

func TestDayArgNotInt(t *testing.T) {
	args := []string{"2015", "notanint"}
	testArgsErrHelper(t, args, DayArgNotInt)
}

func TestYearBelowMin(t *testing.T) {
	args := []string{"2014", "4"}
	testArgsErrHelper(t, args, YearArgInvalid)
}

func TestYearAboveMax(t *testing.T) {
	testYear := fmt.Sprintf("%d", time.Now().Year()+2)
	args := []string{testYear, "4"}
	testArgsErrHelper(t, args, YearArgInvalid)
}

func TestDayBelowMin(t *testing.T) {
	args := []string{"2015", "0"}
	testArgsErrHelper(t, args, DayArgInvalid)
}

func TestDayAboveMaxPastYear(t *testing.T) {
	args := []string{"2015", "26"}
	testArgsErrHelper(t, args, DayArgInvalid)
}

// TODO: Test for current year before AoC starts and while it is running
