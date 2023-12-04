package main

import (
	"fmt"
	"regexp"
	"strconv"
	"util/util"
)

func main() {
	lines, _ := util.ReadFileLineByLine("./day-1/input.txt")
	sum := 0
	nums := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9)`)
	for _, line := range lines {
		fmt.Println(line)
		matches := re.FindAllString(line, -1)
		if matches == nil {
			fmt.Println("No matches found")
		}
		fmt.Println(matches)
		for x, match := range matches {
			for i, num := range nums {
				if match == num {
					matches[x] = strconv.Itoa(i)
				}
			}
		}
		fmt.Println(matches)

		firstLast := matches[0] + matches[len(matches)-1]
		intValue, err := strconv.Atoi(firstLast)
		if err == nil {
			sum += intValue
		}
		fmt.Println(intValue)
		fmt.Println(sum)

	}
	fmt.Println(sum)
}
