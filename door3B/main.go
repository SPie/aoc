package main

import (
	"advent/diagnostic"
	"advent/input"
	"fmt"
	"log"
	"strings"
)

type node struct {
	left *node
	right *node
	values []string
}

func newNode() *node {
	return &node{values: []string{}}
}

func (n *node) setLeft(leftNode *node) {
	n.left = leftNode
}

func (n node) getLeft() *node {
	return n.left
}

func (n *node) setRight(rightNode *node) {
	n.right = rightNode
}

func (n node) getRight() *node {
	return n.right
}

func (n *node) addValue(value string) {
	n.values = append(n.values, value)
}

func (n node) count() int {
	leftCount := 0
	left := n.getLeft()
	if left != nil {
		leftCount = left.count()
	}

	rightCount := 0
	right := n.getRight()
	if right != nil {
		rightCount = right.count()
	}

	return leftCount + rightCount + len(n.values)
}

func main() {
	lines, err := input.GetInput("input")
	if err != nil {
		log.Fatal(err)
	}

	root := buildTree(lines)

	current := root
	for len(current.values) == 0 {
		current = getNextOxygenGeneratorNode(current)
	}

	oxygenGeneratorRating := current.values[0]

	current = root
	for len(current.values) == 0 {
		current = getNextO2ScrubberRating(current)
	}

	co2ScrubberRating := current.values[0]

	lifeSupportRating, err := diagnostic.CalculateRateFromBinaries(oxygenGeneratorRating, co2ScrubberRating)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lifeSupportRating)
}

func buildTree(lines []string) *node {
	root := newNode()
	for _, binary := range lines {
		current := root
		for _, bit := range strings.Split(binary, "") {
			current = getNextNode(current, bit)
		}

		current.addValue(binary)
	}

	return root
}

func getNextNode(current *node, bit string) *node {
	if bit == "1" {
		left := current.getLeft()
		if left == nil {
			left = newNode()
			current.setLeft(left)
		}

		return left
	} 

	right := current.getRight()
	if right == nil {
		right = newNode()
		current.setRight(right)
	}

	return right
}

func getNextOxygenGeneratorNode(current *node) *node {
	left := current.getLeft()
	if left == nil {
		return current.getRight()
	}

	right := current.getRight()
	if right == nil {
		return left
	}

	if left.count() >= right.count() {
		return left
	}

	return right
}

func getNextO2ScrubberRating(current *node) *node {
	left := current.getLeft()
	if left == nil {
		return current.getRight()
	}

	right := current.getRight()
	if right == nil {
		return left
	}

	if left.count() < right.count() {
		return left
	}

	return right
}