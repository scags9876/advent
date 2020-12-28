package main

import (
	"fmt"
)

type publicKeys struct {
	card int
	door int
}

var input = publicKeys{
	card: 1614360,
	door: 7734663,
}
var testInput = publicKeys{
	card: 5764801,
	door: 17807724,
}

const (
	startValue = 1
	divisor    = 20201227
	maxLoops   = 10000000
)

const verbose = true

func main() {
	fmt.Println("start")
	result := part1(testInput)
	if result == 14897079 {
		fmt.Printf("Test input success!\n")
	} else {
		panic("nope")
	}
	result = part1(input)
	fmt.Printf("Part 1: encryption key is %d\n", result)
}

func part1(keys publicKeys) int {
	fmt.Printf("\nStarting part 1 with door key %d and card key %d\n", keys.door, keys.card)

	doorLoops := calcLoopsToPublicKey(keys.door, 7)
	cardLoops := calcLoopsToPublicKey(keys.card, 7)

	if verbose {
		fmt.Printf("Door uses %d loops and card uses %d loops\n", doorLoops, cardLoops)
	}
	// The card transforms the subject number of the door's public key according to the card's
	// loop size. The result is the encryption key.
	cardKey := transformKey(keys.door, cardLoops)

	// The door transforms the subject number of the card's public key according to the door's
	// loop size. The result is the same encryption key as the card calculated.
	doorKey := transformKey(keys.card, doorLoops)

	var encryptionKey int
	if cardKey == doorKey {
		encryptionKey = cardKey
	}

	fmt.Printf("Encryption key is %d\n\n", encryptionKey)
	return encryptionKey
}

func calcLoopsToPublicKey(key, subjectNumber int) (loops int) {
	value := startValue
	for {
		if value == key {
			if verbose {
				fmt.Printf("%d == %d  Success!\n", value, key)
			}
			break
		}
		if loops > maxLoops {
			panic("something might not be right")
		}
		loops++
		// Set the value to itself multiplied by the subject number.
		value = value * subjectNumber
		// Set the value to the remainder after dividing the value by 20201227.
		value = value % divisor
		if verbose {
			//fmt.Printf("After %d loops, looking for %d, value is %d\n", loops, key, value)
		}
	}
	return loops
}

func transformKey(subjectNumber int, loops int) int {
	value := startValue
	for i := 1; i <= loops; i++ {
		// Set the value to itself multiplied by the subject number.
		value = value * subjectNumber
		// Set the value to the remainder after dividing the value by 20201227.
		value = value % divisor
		if verbose {
			//fmt.Printf("After %d of %d loops, value is %d\n", i, loops, value)
		}
	}
	return value
}
