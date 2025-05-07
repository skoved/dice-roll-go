// Copyright skoved
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Represent a roll of the dice. Example: 2d6
type roll struct {
	numDice int64
	sides   int64
}

// Calculate the roll of r
func (r roll) Roll() *big.Int {
	var i int64 = 0
	sum := big.NewInt(0)
	for ; i < r.numDice; i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(r.sides))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not roll the dice\nsides: %d, number of dice: %d\nError: %s\n", r.sides, r.numDice, err)
			os.Exit(1)
		}
		// range for val is 0 to r.sides-1 but we want 1 to r.sides
		val.Add(val, big.NewInt(1))
		sum.Add(sum, val)
	}
	return sum
}

// Return the ROLL(s) from stdin
func rollFromStdin() []roll {
	rolls := []roll{}
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read input from stdin:", err)
	}

	var (
		rollBuf   bytes.Buffer
		isNumDice = true
		tmpRoll   roll
	)
	for i, b := range input {
		char := rune(b)
		if unicode.IsDigit(char) {
			rollBuf.WriteByte(b)
		} else if unicode.IsSpace(char) {
			if isNumDice {
				tmpRoll.numDice, err = strconv.ParseInt(rollBuf.String(), 10, 64)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Could not read dice roll", i, err)
					os.Exit(1)
				}
				isNumDice = false
				rollBuf.Reset()
			} else {
				tmpRoll.sides, err = strconv.ParseInt(rollBuf.String(), 10, 64)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Could not read dice roll", i, err)
					os.Exit(1)
				}
				// toggle between numDice and sidesBuf
				isNumDice = true
				rollBuf.Reset()
				rolls = append(rolls, tmpRoll)
			}
		} else if char == 'd' {
			// d can only come after numDice
			if !isNumDice {
				fmt.Fprintln(os.Stderr, "Invalid ROLL. Check 'd' at position", i)
				os.Exit(1)
			}
			tmpRoll.numDice, err = strconv.ParseInt(rollBuf.String(), 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not read dice roll", i, err)
				os.Exit(1)
			}
			isNumDice = false
			rollBuf.Reset()
		} else {
			fmt.Fprintln(os.Stderr, "Character", char, "is not allowed in a ROLL at position", i)
			os.Exit(1)
		}
	}

	return rolls
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
			fmt.Fprintln(os.Stderr, "Could not read dice roll", i, err)
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
			fmt.Fprintln(os.Stderr, "Could not read dice roll", i, err)
			os.Exit(1)
		}

		rolls = append(rolls, roll{
			numDice: numDice,
			sides:   sides,
		})
	}
	return rolls
}
