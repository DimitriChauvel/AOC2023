package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	partOne := 0
	partTwo := 0
	
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	regex, _ := regexp.Compile("([0-9]+) (red|green|blue)")	
	
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindAllStringSubmatch(line, -1)
		
		isPossible := true
		bag := NewMinBag()

		for _, value := range match {
			color := value[2]
			number, _ := strconv.Atoi(value[1])

			bag.SetMinBag(color, number)
			IdentifyGameIsPossible(&isPossible, color, number)
		}

		index++

		partTwo += bag.Power()

		if !isPossible {
			continue
		}
		partOne += index
	}

	fmt.Printf("Part One: %d\n", partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}
