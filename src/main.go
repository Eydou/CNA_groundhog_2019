//
// EPITECH PROJECT, 2020
// CNA_groundhog_2019
// File description:
// main
//

package main

import (
	"fmt"
	"os"
	"strconv"

	functions "./functions"
)

func help() {
	fmt.Printf("SYNOPSIS\n")
	fmt.Printf("\t./groundhog period\n\n")
	fmt.Printf("DESCRIPTION\n")
	fmt.Printf("\tperiod\tthe number of days defining a period\nt")
	os.Exit(0)
}

func main() {
	args := os.Args

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			help()
		}
		if _, err := functions.ErrorArgs(args); err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mX\033[0m Error: %s\n", err)
			os.Exit(84)
		}
		number1, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err, number1)
			os.Exit(84)
		}
		if number1 == 0 {
			os.Exit(84)
		}
		functions.GroundHog(number1)
		os.Exit(0)
	}
	os.Exit(84)
}
