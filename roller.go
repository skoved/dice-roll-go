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

func dropRollerOpt(drops uint) rolleropt {
	return func(r *roller) {
		r.numDrops = drops
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
	numDrops uint
}

func newRoller(rolls []roll, opts ...rolleropt) roller {
	r := roller{
		print:    defaultPrinter,
		rolls:    rolls,
		numRolls: 1,
		numDrops: 0,
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
			r.print(roll.Roll(dropRollOpt(int(r.numDrops))))
		}
	}
}
