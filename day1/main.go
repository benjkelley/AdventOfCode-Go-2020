package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check_sum(base int, nums []int) {
	for _, num := range nums {
		if num+base == 2020 {
			fmt.Printf("%d %d == 2020\n", base, num)
			sum := base * num
			fmt.Printf("%d == sum\n", sum)
			return
		}
	}
	return
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

	for _, num := range nums {
		check_sum(num, nums)

	}

}
