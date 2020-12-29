package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* take in the first number to compare, the entire int slice, and the recursion level
* if we're at level 1 of recursion, pass the sum of the given num with each element and call this func again
* once we have a set that sums to 2020, perform the multiplication and print to the screen
 */
func check_sum(base int, nums []int, level int) int {
	for _, num := range nums {
		if level == 1 {
			level2 := level + 1
			base3 := check_sum(num+base, nums, level2)
			if base3 != 0 {
				fmt.Printf("%d %d %d == 2020\n", base, num, base3)
				sum := base * num * base3
				fmt.Printf("%d == sum\n", sum)
			}
		}
		if num+base == 2020 {
			return num
		}
	}
	return 0
}

/* 	open the given input file and split the lines into an int slice
* 	start the check_sum recursion
 */
func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)
	var nums []int

	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		nums = append(nums, num)
	}

	level := 1

	for _, num := range nums {
		check_sum(num, nums, level)

	}

}
