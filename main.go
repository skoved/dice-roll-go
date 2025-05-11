// Copyright skoved
// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultNumDrop   = 0
	dropUsage        = "the number of dice from your roll that should be dropped"
	defaultNumRepeat = 1
	repeatUsage      = "the number of times to repeat your dice roll"
	defaultHelp      = false
	helpUsage        = "print help info"
)

var (
	numDrop   int
	numRepeat uint
	help      bool
)

// create command flags
func init() {
	flag.IntVar(&numDrop, "drop", defaultNumDrop, dropUsage)
	flag.IntVar(&numDrop, "d", defaultNumDrop, dropUsage+"(shorthand)")
	flag.UintVar(&numRepeat, "repeat", defaultNumRepeat, repeatUsage)
	flag.UintVar(&numRepeat, "r", defaultNumRepeat, repeatUsage+"(shorthand)")
	flag.BoolVar(&help, "help", defaultHelp, helpUsage)
	flag.BoolVar(&help, "h", defaultHelp, helpUsage+"(shorthand)")
}

// Prints the help/usage info and exits with status code 0
func helpInfo() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS]... [ROLL]\n", os.Args[0])
	fmt.Fprintln(flag.CommandLine.Output(), "Roll the dice from ROLL.")
	fmt.Fprintln(flag.CommandLine.Output(), "Example: roll 2d6")
	fmt.Fprintln(flag.CommandLine.Output(), "Rolls a six sided dice two times and adds the results together.")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	flag.PrintDefaults()
	os.Exit(0)
}

// If the drop flag has a negative value, exits with code 1
func ValidateFlagValues() {
	if numDrop < 0 {
		fmt.Fprintln(os.Stderr, "--drop cannot be negative number")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	if help {
		helpInfo()
	}

	roller := newRoller(getRolls(), repeatRollerOpt(numRepeat), dropRollerOpt(numDrop))
	roller.roll()

	fmt.Fprintln(os.Stderr, "drop:", numDrop)
	fmt.Fprintln(os.Stderr, "repeat:", numRepeat)
}
