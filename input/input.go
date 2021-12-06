package input

import (
	"io/ioutil"
	"strings"
)

func GetInput(fileName string) ([]string, error) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(input), "\n"), nil
}