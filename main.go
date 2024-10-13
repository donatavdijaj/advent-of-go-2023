package main

import (
	day1 "advent-of-go-2023/day-1/part-1"
	"fmt"
)

func main() {
	sum, err := day1.Calibrate("day-1/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Day 1: %d\n", sum)
}
