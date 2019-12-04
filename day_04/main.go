package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLines(filename string) []string {
	file, err := os.Open("inputs/" + filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func check(number int) bool {
	double := false
	for number > 0 {
		last := number % 10
		pen := (number / 10) % 10
		if last < pen {
			return false
		}
		if last == pen {
			double = true
		}
		number = number / 10
	}
	return double
}

func checkNoGroups(number int) bool {
	var repeats [10]int
	for number > 0 {
		last := number % 10
		pen := (number / 10) % 10
		if last < pen {
			return false
		}
		if last == pen {
			if repeats[last] == 0 {
				repeats[last]++
			}
			repeats[last]++
		}
		number = number / 10
	}

	double := false
	for i := range repeats {
		if repeats[i] == 2 {
			double = true
		}
	}
	return double
}

func partA(input string) string {
	count := 0
	split := strings.Split(getLines(input)[0], "-")
	start, _ := strconv.Atoi(split[0])
	finish, _ := strconv.Atoi(split[1])

	for ; start < finish; start++ {
		if check(start) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func partB(input string) string {
	count := 0
	split := strings.Split(getLines(input)[0], "-")
	start, _ := strconv.Atoi(split[0])
	finish, _ := strconv.Atoi(split[1])

	for ; start < finish; start++ {
		if checkNoGroups(start) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
