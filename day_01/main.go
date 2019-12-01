package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func calculateFuel(f int) int {
	fuel := f/3 - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + calculateFuel(fuel)
}

func partA(input string) string {
	total := 0
	for _, line := range getLines(input) {
		mass, _ := strconv.Atoi(line)
		total += (mass/3 - 2)
	}
	return strconv.Itoa(total)
}

func partB(input string) string {
	total := 0
	for _, line := range getLines(input) {
		mass, _ := strconv.Atoi(line)
		total += calculateFuel(mass)
	}
	return strconv.Itoa(total)
}

func main() {
	input := "real.txt"
	fmt.Println("A: " + partA(input))
	fmt.Println("B: " + partB(input))
}
