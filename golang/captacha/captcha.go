package main

import (
	"fmt"
	"os"
	"strconv"
)

func getSum(arr []string) int {
	sum := 0
	for i, num := range arr {
		current, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("err", err)
		}
		var next int
		if i+1 == len(arr) {
			next, err = strconv.Atoi(arr[0])
		} else {
			next, err = strconv.Atoi(arr[i+1])
		}

		if current == next {
			sum = sum + current
		}
	}
	return sum
}

func main() {
	arr := os.Args[1:]
	sum := getSum(arr)
	fmt.Println("sum", sum)
}
