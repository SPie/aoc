package main

import (
	"advent/input"
	"fmt"
	"log"
	"strconv"
)

func main() {
	lines, err := input.GetInput("input")
	if err != nil {
		log.Fatal(err)
	}

	increased := 0
	for i := 3; i < len(lines); i++ {
		first, err := sumMeasurementGroup(lines[i - 3], lines[i - 2], lines[i - 1])
		if err != nil {
			log.Fatal(err)
		}
		second, err := sumMeasurementGroup(lines[i - 2], lines[i - 1], lines[i])
		if err != nil {
			log.Fatal(err)
		}

		if second > first {
			increased++
		}
	}

	fmt.Println(increased)
}

func sumMeasurementGroup(m1 string, m2 string, m3 string) (int, error) {
	m1int, err := strconv.Atoi(m1)
	if err != nil {
		return 0, err
	}
	m2int, err := strconv.Atoi(m2)
	if err != nil {
		return 0, err
	}
	m3int, err := strconv.Atoi(m3)
	if err != nil {
		return 0, err
	}

	return m1int + m2int + m3int, nil
}
