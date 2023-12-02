package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("document")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int

	for scanner.Scan() {
		text := scanner.Text()
		nums := ""

		for _, v := range text {
			if _, err := strconv.Atoi(string(v)); err == nil {
				nums += string(v)
			}
		}

		nums = string(nums[0]) + string(nums[len(nums)-1])

		n, _ := strconv.Atoi(nums)
		sum += n
	}

	fmt.Println(sum)
}
