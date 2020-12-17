package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"

func main() {
	arrivalTime, busSchedule := GetInput(inputFilename)
	part1(arrivalTime, busSchedule)
	part2(busSchedule)
}

func GetInput(inputFilename string) (int, []int) {
	input := lib.GetInputStrings(inputFilename)

	arrivalTime, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	var busSchedule []int
	for _, n := range strings.Split(input[1], ",") {
		busID := 0
		if n != "x" {
			busID, err = strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
		}
		busSchedule = append(busSchedule, busID)
	}

	return arrivalTime, busSchedule
}

func part1(arrivalTime int, busSchedule []int) {
	result := 0
	fmt.Printf("arrivalTime: %d, busSchedule: %v\n", arrivalTime, busSchedule)

	var earliestBus, earliestBusArrival int
	for _, busID := range busSchedule {
		if busID == 0 {
			continue
		}
		busArrivalTime := busID * int(math.Ceil(float64(arrivalTime)/float64(busID)))
		fmt.Printf("aT/bID=%d/%d=%.02f\n", arrivalTime, busID, float64(arrivalTime)/float64(busID))
		fmt.Printf("bAT=bID*ceil(aT/BID)=%d*%d=%d\n", busID, int(math.Ceil(float64(arrivalTime)/float64(busID))), busArrivalTime)
		if earliestBusArrival == 0 || busArrivalTime < earliestBusArrival {
			earliestBusArrival = busArrivalTime
			earliestBus = busID
		}
	}

	waitTime := earliestBusArrival - arrivalTime
	result = earliestBus * waitTime

	fmt.Printf("Part 1: %d\n", result)
}

func part2(busSchedule []int) {
	offsets := make(map[int]int)
	sortedSchedule := make([]int, 0)
	for i, busID := range busSchedule {
		if busID == 0 {
			continue
		}
		offsets[busID] = i
		sortedSchedule = lib.SortedInsertInt(sortedSchedule, busID)

		if big.NewInt(int64(busID)).ProbablyPrime(0) {
			fmt.Println(busID, "is prime")
		} else {
			fmt.Println(busID, "is not prime")
		}
	}

	// start with the largest buses first time around, guaranteeing we will start skipping
	// by this much at the outset of the algorithm
	largestBusID := sortedSchedule[len(sortedSchedule)-1]
	timestamp := largestBusID - offsets[largestBusID]

	done := false
	for !done {
		fmt.Printf("[%10d] ", timestamp)
		timestampSkip := 1
		foundMagicTimestamp := true
		for i := len(sortedSchedule) - 1; i >= 0; i-- {
			busID := sortedSchedule[i]
			stepValue := timestamp + offsets[busID]

			fmt.Printf("%d/%d=%dr%d ", stepValue, busID, stepValue/busID, stepValue%busID)

			if stepValue%busID != 0 {
				fmt.Printf("Nope! skipping ahead by %d\n", timestampSkip)
				foundMagicTimestamp = false
				break
			}
			// This bus matches, so skip ahead by this bus's ID times the current skip
			// when 2 or more buses are in sync, we know this wont happen again for another
			// x minutes where x = busAID * busBID
			// in order for this to work, it's important that these busIDs are all prime numbers
			timestampSkip *= busID
		}

		if foundMagicTimestamp {
			fmt.Printf("Found it!\n")
			done = true
			break
		}
		//if step == 10 {
		//	done = true
		//}
		timestamp += timestampSkip
	}

	fmt.Printf("Part 2: %d\n", timestamp)
}
