// Copyright skoved
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"math/big"
)

type rolleropt func(*roller)

func repeatRollerOpt(repeats uint) rolleropt {
	return func(r *roller) {
		r.numRolls = repeats
	}
}

type printer func(*big.Int)

func defaultPrinter(val *big.Int) {
	fmt.Println(val)
}

type roller struct {
	print    printer
	rolls    []roll
	numRolls uint
}

func newRoller(rolls []roll, opts ...rolleropt) roller {
	r := roller{
		rolls:    rolls,
		numRolls: 1,
		print:    defaultPrinter,
	}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}

// roll the dice from rolls and print the result
func (r roller) roll() {
	for _, roll := range r.rolls {
		for range r.numRolls {
			r.print(roll.Roll())
		}
	}
}
