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

// argModes are enums that represent where the ROLL instruction can be found. 0 means ROLL will be provided in stdin. >=
// 1 means that ROLL was passed as an argument to the command. The value for argMode should be compared to flag.NArg().
type argMode int

const stdin argMode = 0

// Returns the argMode for the invocation of roll.
func getArgMode() argMode {
	return argMode(flag.NArg())
}

// Returns the ROLL. It detects if the ROLL was provided as an arg or if it is passed through stdin
func (a argMode) getRoll() string {
	switch a {
	case stdin:
		return rollFromStdin()
	default:
		return rollFromArg()
	}
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

// Return the ROLL from stdin
func rollFromStdin() string {
	fmt.Println("Need to implement reading roll from stdin")
	return ""
}

func rollFromArg() string {
	fmt.Println("Need to implement parsing roll from cmd arg")
	return ""
}

func main() {
	flag.Parse()

	if help {
		helpInfo()
	}

	mode := getArgMode()
	mode.getRoll()

	fmt.Println("drop:", numDrop)
	fmt.Println("repeat:", numRepeat)
}
