package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/al-ce/aoc-go-fetch/fetchInput"
	"github.com/al-ce/aoc-go-fetch/validateArgs"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, ".env file not found or couldn't be loaded: %v", err)
		fmt.Fprintf(os.Stderr, "Add your AoC session cookie to the .env file")
		os.Exit(1)
	}
}

func main() {
	loadEnv()

	verbose := flag.Bool("v", false, "verbose output")

	flag.Parse()

	args := flag.Args()

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

	input, err := fetchInput.GetPuzzleInput(year, day)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nFetch Error: %s\n", err)
		return
	}

	fmt.Println("🪚 input", input)
}
