package setup

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (string, error) {
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

func InitFlags() {
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

func GetArgs() []string {
	return flag.Args()
}
