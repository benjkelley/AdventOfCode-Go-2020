package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func tree_checker(run int, rise int, lines []string) int {
	currentCol := 0
	treeCount := 0

	for i := 0; i < len(lines); i += rise {
		nodes := strings.Split(lines[i], "")
		// since the columns continue in the same pattern infinitely to the right, we can just loop back to the start of the array
		// each time we go out of bounds
		if currentCol >= len(nodes) {
			currentCol -= len(nodes)
		}
		if nodes[currentCol] == "#" {
			treeCount++
		}
		currentCol = currentCol + run
	}

	return treeCount
}

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	case1 := tree_checker(1, 1, lines)
	fmt.Printf("Tree Count for (1, 1): %d\n", case1)

	case2 := tree_checker(3, 1, lines)
	fmt.Printf("Tree Count for (3, 1): %d\n", case2)

	case3 := tree_checker(5, 1, lines)
	fmt.Printf("Tree Count for (5, 1): %d\n", case3)

	case4 := tree_checker(7, 1, lines)
	fmt.Printf("Tree Count for (7, 1): %d\n", case4)

	case5 := tree_checker(1, 2, lines)
	fmt.Printf("Tree Count for (1, 2): %d\n\n", case5)

	fmt.Printf("Total: %d\n", (case1 * case2 * case3 * case4 * case5))

}
