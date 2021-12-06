package main

import (
	"advent/input"
	"advent/submarine"
	"fmt"
	"log"
)

func main() {
	lines, err := input.GetInput("input")
	if err != nil {
		log.Fatal(err)
	}

	sub := submarine.NewAimSubmarine()
	for _, line := range lines {
		command, err := submarine.NewCommandByCommandLine(line)
		if err != nil {
			log.Fatal(err)
		}

		err = sub.Move(command)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(sub.Result())
}