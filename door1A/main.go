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
	for i := 1; i < len(lines); i++ {
		first, err := strconv.Atoi(lines[i - 1])
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatal(err)
		}

		if second > first {
			increased++
		}
	}

	fmt.Println(increased)
}