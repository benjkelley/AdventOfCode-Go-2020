package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)
	var lines []string
	currentCol := 0
	treeCount := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		nodes := strings.Split(line, "")
		// since the columns continue in the same pattern infinitely to the right, we can just loop back to the start of the array
		// each time we go out of bounds
		if currentCol >= len(nodes) {
			currentCol -= len(nodes)
		}
		if nodes[currentCol] == "#" {
			treeCount++
		}
		currentCol = currentCol + 3
	}
	fmt.Printf("Tree Count: %d\n", treeCount)

}
