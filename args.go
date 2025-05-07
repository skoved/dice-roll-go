// Copyright skoved
// SPDX-License-Identifier: MIT

package main

import "flag"

// argModes are enums that represent where the ROLL instruction can be found. 0 means ROLL will be provided in stdin. >=
// 1 means that ROLL was passed as an argument to the command. The value for argMode should be compared to flag.NArg().
type argMode int

const stdin argMode = 0

// Returns the ROLL. It detects if the ROLL was provided as an arg or if it is passed through stdin
func getRoll() []roll {
	if argMode(flag.NArg()) == stdin {
		return rollFromStdin()
	}
	return rollsFromArg()
}
