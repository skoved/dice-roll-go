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
)

var (
	numDrop   uint
	numRepeat uint
	help      bool
)

// create command flags
func init() {
	flag.UintVar(&numDrop, "drop", defaultNumDrop, dropUsage)
	flag.UintVar(&numDrop, "d", defaultNumDrop, dropUsage+"(shorthand)")
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

func main() {
	flag.Parse()

	if help {
		helpInfo()
	}

	roller := newRoller(getRolls(), repeatRollerOpt(numRepeat))
	roller.roll()

	fmt.Fprintln(os.Stderr, "drop:", numDrop)
	fmt.Fprintln(os.Stderr, "repeat:", numRepeat)
}
