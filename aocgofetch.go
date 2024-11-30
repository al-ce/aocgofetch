package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/al-ce/aocgofetch/fetchInput"
	"github.com/al-ce/aocgofetch/validateArgs"
)

func loadEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		err = fmt.Errorf(".env file not found or couldn't be loaded: %v", err)
		return "", err
	}
	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		err := fmt.Errorf("Set AOC_SESSION environment variable")
		return "", err
	}
	return sessionCookie, nil
}

func initFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Advent of Code Input Fetcher\n\n")

		fmt.Fprintf(os.Stderr, "Usage:\n  %s [options] <day> <year>\n\n", os.Args[0])

		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(os.Stderr, "  day\t\tDay of the puzzle (1-25)\n")
		fmt.Fprintf(os.Stderr, "  year\t\tYear of the puzzle (2015-present)\n\n")

		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  Get day 1 input from 2023:\n")
		fmt.Fprintf(os.Stderr, "    $ %s 1 2023\n", os.Args[0])
	}

	flag.Parse()
}

func main() {
	// Ensure user set their session cookie in the .env file
	sessionCookie, err := loadEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nEnvironment Error: %s\n", err)
		os.Exit(1)
	}

	// Set flags and validate args
	initFlags()
	args := flag.Args()
	year, day, argsState := validateArgs.GetYearAndDay(args)

	if argsState != validateArgs.ValidArgs {
		fmt.Fprintf(os.Stderr, "aocgofetch: %s", argsState.Error(args))
		os.Exit(1)
	}

	// Fetch the puzzle input
	input, err := fetchInput.GetPuzzleInput(year, day, sessionCookie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "aocgofetch: %s\n", err)
		os.Exit(1)
	}

	fmt.Print(input)

	os.Exit(0)
}
