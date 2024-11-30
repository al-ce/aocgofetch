package main

import (
	"fmt"
	"os"

	"github.com/al-ce/aocgofetch/fetchInput"
	"github.com/al-ce/aocgofetch/parser"
	"github.com/al-ce/aocgofetch/setup"
)

func main() {
	// Ensure user set their session cookie in the .env file
	sessionCookie, err := setup.LoadEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nEnvironment Error: %s\n", err)
		os.Exit(1)
	}

	// Set flags and validate args
	setup.InitFlags()
	parsed, argsErr := parser.Parse(setup.GetArgs())
	if argsErr != parser.ValidArgs {
		err = parsed.FmtArgsErr(argsErr)
		fmt.Fprintf(os.Stderr, "aocgofetch: %s", err)
		os.Exit(1)
	}

	// Fetch the puzzle input
	input, err := fetchInput.GetPuzzleInput(parsed.Year, parsed.Day, sessionCookie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "aocgofetch: %s\n", err)
		os.Exit(1)
	}

	fmt.Print(input)

	os.Exit(0)
}
