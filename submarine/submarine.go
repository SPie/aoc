package submarine

import (
	"errors"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	fields int
}

func NewCommandByCommandLine(line string) (Command, error) {
	commandLine := strings.Split(line, " ")
	fields, err := strconv.Atoi(commandLine[1])
	if err != nil {
		return Command{}, err
	}

	return NewCommand(commandLine[0], fields), nil

}

func NewCommand(direction string, fields int) Command {
	return Command{direction: direction, fields: fields}
}

type Submarine interface {
	Move(command Command) error
	Result() int
}

type classicSubmarine struct {
	x int
	y int
}

func NewSubmarine() *classicSubmarine {
	return &classicSubmarine{x: 0, y: 0}
}

func (s *classicSubmarine) Move(command Command) error {
	switch command.direction {
	case "forward":
		s.y = s.y + command.fields
	case "down":
		s.x += command.fields
	case "up":
		s.x -= command.fields
	default:
		return errors.New("Invalid direction")
	}	

	return nil
}

func (s classicSubmarine) Result() int {
	return s.x * s.y
}

type aimSubmarine struct {
	x int
	y int
	aim int
}

func NewAimSubmarine() *aimSubmarine {
	return &aimSubmarine{x: 0, y: 0, aim: 0}
}

func (as *aimSubmarine) Move(command Command) error {
	switch command.direction {
	case "forward":
		as.y += command.fields
		as.x += (as.aim * command.fields)
	case "down":
		as.aim += command.fields
	case "up":
		as.aim -= command.fields
	default:
		return errors.New("Invalid direction")
	}

	return nil
}

func (as aimSubmarine) Result() int {
	return as.x * as.y
}