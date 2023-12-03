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
	nums := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	re := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9|0)`)

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		fmt.Println(matches)
		for x, match := range matches {
			for i, num := range nums {
				if match == num {
					fmt.Println(i)
					matches[x] = string(i)
				}
			}
		}
		fmt.Println(matches)

		firstLast := matches[0] + matches[len(matches)-1]
		intValue, err := strconv.Atoi(firstLast)
		if err == nil {
			sum += intValue
		}

	}
	fmt.Println(sum)
}
