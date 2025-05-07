// Copyright skoved
// SPDX-License-Identifier: MIT

package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type roll struct {
	numDice int64
	sides   int64
}

func (r roll) Roll() *big.Int {
	var i int64 = 0
	sum := big.NewInt(0)
	for ; i < r.numDice; i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(r.sides))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not roll the dice\nsides: %d, number of dice: %d\nError: %s\n", r.sides, r.numDice, err)
			os.Exit(1)
		}
		// range for val is 0 to r.sides-1 but we want 1 - r.sides
		val.Add(val, big.NewInt(1))
		sum.Add(sum, val)
	}
	return sum
}

// Return the ROLL from stdin
func rollFromStdin() []roll {
	fmt.Println("Need to implement reading roll from stdin")
	return []roll{}
}

// Return the ROLL(s) from the command line args
func rollsFromArg() []roll {
	args := flag.Args()
	rolls := []roll{}
	for i, arg := range args {
		var (
			j   int
			num strings.Builder
		)
		for ; j < len(arg) && arg[j] != 'd'; j++ {
			num.WriteByte(arg[j])
		}
		numDice, err := strconv.ParseInt(num.String(), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read dice roll ", i, err)
			os.Exit(1)
		}

		// increment past d
		j++

		num.Reset()
		for ; j < len(arg) && arg[j] != 'd'; j++ {
			num.WriteByte(arg[j])
		}
		sides, err := strconv.ParseInt(num.String(), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read dice roll ", i, err)
			os.Exit(1)
		}

		rolls = append(rolls, roll{
			numDice: numDice,
			sides:   sides,
		})
	}
	return rolls
}
