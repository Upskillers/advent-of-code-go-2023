package main

import (
	"fmt"
	"strconv"
	"unicode"
	"util/util"
)

func main() {
	lines, _ := util.ReadFileLineByLine("./day-1/input.txt")
	sum := 0

	for _, line := range lines {
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

	}
	fmt.Println(sum)

}
