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
	defaultNumRepeat = 0
	repeatUsage      = "the number of times to repeat your dice roll"
	defaultHelp      = false
	helpUsage        = "print help info"
	defaultVersionF  = false
	helpVersionF     = "print version"
)

var (
	numDrop   int
	numRepeat uint
	help      bool
	versionF  bool
	version   string
)

// create command flags
func init() {
	flag.IntVar(&numDrop, "drop", defaultNumDrop, dropUsage)
	flag.IntVar(&numDrop, "d", defaultNumDrop, dropUsage+"(shorthand)")
	flag.UintVar(&numRepeat, "repeat", defaultNumRepeat, repeatUsage)
	flag.UintVar(&numRepeat, "r", defaultNumRepeat, repeatUsage+"(shorthand)")
	flag.BoolVar(&help, "help", defaultHelp, helpUsage)
	flag.BoolVar(&help, "h", defaultHelp, helpUsage+"(shorthand)")
	flag.BoolVar(&versionF, "version", defaultVersionF, helpVersionF)
	flag.BoolVar(&versionF, "v", defaultVersionF, helpVersionF+"(shorthand)")
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
func validateFlagValues() {
	if numDrop < 0 {
		fmt.Fprintln(os.Stderr, "--drop cannot be a negative number")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	if versionF {
		fmt.Println(os.Args[0], "version", version)
		fmt.Println("Copyright (C) 2025 Sam Koved")
		fmt.Println("License MIT <https://github.com/skoved/dice-roll-go/blob/main/LICENSE>")
		fmt.Println()
		fmt.Println("Written by Sam Koved <https://github.com/skoved>")
		os.Exit(0)
	}

	validateFlagValues()

	if help {
		helpInfo()
	}

	roller := newRoller(getRolls(), repeatRollerOpt(numRepeat), dropRollerOpt(numDrop))
	roller.roll()
}
