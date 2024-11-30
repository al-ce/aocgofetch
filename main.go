package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/al-ce/aoc-go-fetch/fetchInput"
	"github.com/al-ce/aoc-go-fetch/validateArgs"
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

func main() {
	sessionCookie, err := loadEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nEnvironment Error: %s\n", err)
		os.Exit(1)
	}

	year, day, argsState := validateArgs.GetYearAndDay(args)

	if argsState != validateArgs.ValidArgs {
		fmt.Fprintf(os.Stderr, "\nArguments Error: %s\n", validateArgs.ArgsErrType[argsState])
		return
	}

	if *verbose {
		fmt.Println("\n\t❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
		fmt.Println("\t❄️Advent of Code Puzzle Input Fetcher ❄️")
		fmt.Println("\t❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
		fmt.Printf("\n\tFetching input for AoC %d day %d...\n", year, day)
	}

	input, err := fetchInput.GetPuzzleInput(year, day, sessionCookie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nFetch Error: %s\n", err)
		return
	}

	fmt.Println("🪚 input", input)
}
