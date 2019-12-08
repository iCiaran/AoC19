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

func buildTreeChildren(input []string) map[string][]string {
	tree := make(map[string][]string)
	for _, s := range input {
		split := strings.Split(s, ")")
		front, back := split[0], split[1]
		tree[front] = append(tree[front], back)
	}
	return tree
}

func buildTreeParents(input []string) map[string]string {
	tree := make(map[string]string)
	for _, s := range input {
		split := strings.Split(s, ")")
		front, back := split[0], split[1]
		tree[back] = front
	}
	return tree
}

func parseTreeChildren(tree map[string][]string) int {
	total := 0
	queue := make(chan string, 50)
	queue <- "COM"
	sum := make(map[string]int, 0)
	sum["COM"] = 0

	for len(queue) > 0 {
		current := <-queue
		total += sum[current]
		for _, o := range tree[current] {
			sum[o] = sum[current] + 1
			queue <- o
		}
	}

	return total
}

func parseTreeParents(tree map[string]string) int {
	path := make(map[string]int, 0)
	current := "YOU"
	count := 0

	for current != "COM" {
		path[tree[current]] = count - 1
		current = tree[current]
		count++
	}

	current = "SAN"
	count = 0

	for current != "COM" {
		if path[current] != 0 {
			return count + path[current]
		}
		current = tree[current]
		count++
	}
	return 0
}

func partA(input string) string {
	tree := buildTreeChildren(getLines(input))
	count := parseTreeChildren(tree)
	return strconv.Itoa(count)
}

func partB(input string) string {
	tree := buildTreeParents(getLines(input))
	count := parseTreeParents(tree)
	return strconv.Itoa(count)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
