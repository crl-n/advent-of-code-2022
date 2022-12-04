package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type assignmentPairRange struct {
	lo int
	hi int
}

func splitAssignmentPair(line string) (string, string) {
	pair := strings.Split(line, ",")
	return pair[0], pair[1]
}

func splitRange(line string) assignmentPairRange {
	pair := strings.Split(line, "-")
	lo, _ := strconv.Atoi(pair[0])
	hi, _ := strconv.Atoi(pair[1])
	return assignmentPairRange{lo, hi}
}

func isContained(first assignmentPairRange, second assignmentPairRange) bool {
	if first.lo >= second.lo && first.hi <= second.hi {
		return true
	}
	return false
}

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	count := 0
	for _, line := range lines {
		first, second := splitAssignmentPair(line)
		
		firstRange := splitRange(first)
		secondRange := splitRange(second)

		if isContained(firstRange, secondRange) || isContained(secondRange, firstRange) {
			count++
		}
	}

	fmt.Println("Count: ", count)
}
