package main

import (
	"flag"
	"fmt"

	"github.com/al-ce/aoc-go-fetch/validateArgs"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")

	flag.Parse()

	args := flag.Args()

	year, day, argsState := validateArgs.GetYearAndDay(args)

	if argsState != validateArgs.ValidArgs {
		fmt.Printf("Arguments Error: %s\n", validateArgs.ArgsErrType[argsState])
		return
	}

	if *verbose {
		fmt.Println("\n\t❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
		fmt.Println("\t❄️Advent of Code Puzzle Input Fetcher ❄️")
		fmt.Println("\t❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️❄️")
		fmt.Printf("\n\tFetching input for AoC %d day %d...\n", day, year)
	}
}
