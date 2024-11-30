package handleErrors

import (
	"fmt"
	"os"

	"github.com/al-ce/aocgofetch/validateArgs"
)

func BadArgs(args []string, argsState validateArgs.ArgsState) {
	badarg := ""
	switch argsState {
	case validateArgs.BadArgsAmount:
		badarg = fmt.Sprintf("got %d args", len(args))
	case validateArgs.YearArgNotInt,
		validateArgs.YearArgInvalid:
		badarg = args[0]
	case validateArgs.DayArgNotInt,
		validateArgs.DayArgInvalid:
		badarg = args[1]
	}
	fmt.Fprintf(os.Stderr, "aocgofetch: %s: %s", badarg, validateArgs.ArgsErrType[argsState])
	os.Exit(1)
}

func FetchError(err error) {
	fmt.Fprintf(os.Stderr, "\naocgofetch: could not fetch: %s\n", err)
	os.Exit(1)
}
