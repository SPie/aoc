package main

import (
	"advent/diagnostic"
	"advent/input"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := input.GetInput("input")
	if err != nil {
		log.Fatal(err)
	}

	bits, err := calculateBitSums(lines)
	if err != nil {
		log.Fatal(err)
	}

	gammaString := make([]string, len(bits))
	epsilonString := make([]string, len(bits))
	threshhold := len(lines) / 2
	for i, bit := range bits {
		gammaString[i] = getGammaBit(bit, threshhold)
		epsilonString[i] = getEpsilonBit(bit, threshhold)
	}

	poserConsumption, err := diagnostic.CalculateRateFromBinaries(strings.Join(gammaString, ""), strings.Join(epsilonString, ""))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(poserConsumption)
}

func calculateBitSums(lines []string) ([]int, error) {
	bits := map[int]int{}
	for _, line := range lines {
		for i, bit := range strings.Split(line, "") {
			current, exists := bits[i]
			if !exists {
				current = 0
			}
			bitNumber, err := strconv.Atoi(string(bit))
			if err != nil {
				return []int{}, err
			}

			bits[i] = current + bitNumber
		}
	}

	bitSlice := make([]int, len(bits))
	for i, bitSum := range bits {
		bitSlice[i] = bitSum
	}

	return bitSlice, nil
}

func getGammaBit(sum int, threshhold int) string {
	if sum >= threshhold {
		return "1"
	}

	return "0"
}

func getEpsilonBit(sum int, threshhold int) string {
	if sum >= threshhold {
		return "0"
	}

	return "1"
}