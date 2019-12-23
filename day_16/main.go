package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var patterns map[int][]int

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

func getList(line string) (result []int) {
	result = make([]int, len(line))
	for i, c := range line {
		result[i] = int(c - '0')
	}
	return
}

func getPattern(layer int, length int) []int {
	if len(patterns[layer]) > 0 {
		return patterns[layer]
	}

	pattern := []int{0, 1, 0, -1}
	result := make([]int, length)
	pos := 1
	count := 0
	for i := 0; i < length; i++ {
		if i < layer {
			result[i] = 0
		} else {
			result[i] = pattern[pos]
			count++
			if count%(layer+1) == 0 {
				pos = (pos + 1) % 4
			}
		}
	}
	patterns[layer] = result
	return result
}

func lastDigit(in int) int {
	digit := in % 10
	if digit < 0 {
		return digit * -1
	}
	return digit
}

func applyPhase(signal []int) []int {
	newSignal := make([]int, len(signal))
	var total int
	for i := 0; i < len(signal); i++ {
		total = 0
		pattern := getPattern(i, len(signal))
		for j := i; j < len(signal); j++ {
			total += signal[j] * pattern[j]
		}
		newSignal[i] = lastDigit(total)
	}
	return newSignal
}

func partA(input string) string {
	signal := getList(getLines(input)[0])
	patterns = make(map[int][]int, 0)
	for i := 0; i < 100; i++ {
		signal = applyPhase(signal)
	}
	return strings.Trim(strings.Replace(fmt.Sprint(signal), " ", "", -1), "[]")[:8]
}

func partB(input string) string {
	in := getLines(input)[0]
	offset, _ := strconv.Atoi(in[:7])
	signal := getList(strings.Repeat(in, 10000))[offset:]

	for i := 0; i < 100; i++ {
		sum := 0
		for index := len(signal) - 1; index >= 0; index-- {
			sum += signal[index]
			signal[index] = lastDigit(sum)
		}
	}

	return strings.Trim(strings.Replace(fmt.Sprint(signal), " ", "", -1), "[]")[:8]
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
