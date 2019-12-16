package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type chemical struct {
	name   string
	amount int
}

type recipe struct {
	output chemical
	input  []chemical
}

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

func getOreRequired(want string, amount int, recipes map[string]recipe, excess map[string]int) int {
	if want == "ORE" {
		return amount
	}

	if excess[want] >= amount {
		excess[want] -= amount
		return 0
	}

	if excess[want] > 0 {
		amount -= excess[want]
		excess[want] = 0
	}

	recipe := recipes[want]
	amountPerbatch := recipe.output.amount

	batchesRequired := amount / amountPerbatch

	if amount%amountPerbatch > 0 {
		batchesRequired++
	}

	ore := 0
	for _, i := range recipe.input {
		ore += getOreRequired(i.name, i.amount*batchesRequired, recipes, excess)
	}

	excess[want] += recipe.output.amount*batchesRequired - amount

	return ore
}

func getChemical(in []string) chemical {
	amount, _ := strconv.Atoi(in[0])
	return chemical{strings.Trim(in[1], ","), amount}
}

func getRecipes(input string) map[string]recipe {
	recipes := make(map[string]recipe, 0)
	for _, line := range getLines(input) {
		split := strings.Split(line, "=>")
		output := getChemical(strings.Split(strings.TrimSpace(split[1]), " "))
		input := make([]chemical, 0)
		for i, s := 0, strings.Split(strings.TrimSpace(split[0]), " "); i < len(s); i += 2 {
			input = append(input, getChemical([]string{s[i], s[i+1]}))
		}
		recipes[output.name] = recipe{output, input}
	}
	return recipes
}

func partA(input string) string {
	recipes := getRecipes(input)
	ore := getOreRequired("FUEL", 1, recipes, make(map[string]int))

	return strconv.Itoa(ore)
}

func partB(input string) string {
	recipes := getRecipes(input)
	target := 1000000000000
	start := 0

	end := target
	for (end-start)/2 > 0 {

		guess := start + (end-start)/2

		result := getOreRequired("FUEL", guess, recipes, make(map[string]int))
		if result > target {
			end = guess
		} else {
			start = guess
		}
	}
	return strconv.Itoa(start)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
