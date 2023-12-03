package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		var nums []string
		for _, char := range line {
			if unicode.IsDigit(char) {
				nums = append(nums, string(char))
			}
		}
		firstLast := nums[0] + nums[len(nums)-1]
		intValue, err := strconv.Atoi(firstLast)
		if err == nil {
			sum += intValue
		}
		fmt.Println(nums)
		fmt.Println(firstLast)
		fmt.Println(line)

	}
	fmt.Println(sum)

}
