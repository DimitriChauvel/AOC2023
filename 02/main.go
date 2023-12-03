package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	re, _ := regexp.Compile("([0-9]+) (red|green|blue)")

	scanner := bufio.NewScanner(f)
	index := 0
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		match := re.FindAllStringSubmatch(text, -1)
		isPossible := true
		
		for _, v := range match {
			n, _ := strconv.Atoi(v[1])

			switch v[2] {
				case "red":
					if n > 12 {
						isPossible = false
					}
				case "green":
					if n > 13 {
						isPossible = false
					}
				case "blue":
					if n > 14 {
						isPossible = false
					}
			}
		}

		index++
		if !isPossible {
			continue
		}

		sum += index
	}

	fmt.Println(sum)
}