package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction byte
	distance  int
}

type coord struct {
	x int
	y int
}

func getInstructions(input string) []instruction {
	var instructions []instruction
	for _, i := range splitString(input) {
		dir := i[0]
		dis, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		instructions = append(instructions, instruction{dir, dis})
	}
	return instructions
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

func splitString(input string) []string {
	tmp := strings.Split(input, ",")
	var values []string

	for _, s := range tmp {
		values = append(values, s)
	}
	return values
}

func addFirstPath(instructions []instruction) map[coord]int {
	set := make(map[coord]int)
	x, y, count := 0, 0, 0
	for _, i := range instructions {
		for n := 0; n < i.distance; n++ {
			count++
			switch i.direction {
			case 'U':
				y += 1
			case 'D':
				y -= 1
			case 'R':
				x += 1
			case 'L':
				x -= 1
			}
			c := coord{x, y}
			if set[c] == 0 {
				set[c] = count
			}
		}
	}
	return set
}

func addSecondPath(instructions []instruction, firstPath map[coord]int) ([]coord, []int) {
	var collisions []coord
	var timedCollisions []int
	x, y, count := 0, 0, 0
	for _, i := range instructions {
		for n := 0; n < i.distance; n++ {
			count++
			switch i.direction {
			case 'U':
				y += 1
			case 'D':
				y -= 1
			case 'R':
				x += 1
			case 'L':
				x -= 1
			}
			c := coord{x, y}
			if firstPath[c] != 0 {
				collisions = append(collisions, c)
				timedCollisions = append(timedCollisions, firstPath[c]+count)
			}
		}
	}
	return collisions, timedCollisions
}

func manhattanDistance(a coord, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func getShortestManhattan(collisions []coord) int {
	min := math.MaxInt32

	for _, c := range collisions {
		d := manhattanDistance(c, coord{0, 0})
		if d < min {
			min = d
		}
	}
	return min
}

func getShortestCollision(collisions []int) int {
	min := math.MaxInt32
	for _, c := range collisions {
		if c < min {
			min = c
		}
	}
	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func partA(input string) string {
	lines := getLines(input)
	pathA := getInstructions(lines[0])
	pathB := getInstructions(lines[1])

	pathACoords := addFirstPath(pathA)
	collisions, _ := addSecondPath(pathB, pathACoords)
	shortestManhattan := getShortestManhattan(collisions)
	return strconv.Itoa(shortestManhattan)
}

func partB(input string) string {
	lines := getLines(input)
	pathA := getInstructions(lines[0])
	pathB := getInstructions(lines[1])

	pathACoords := addFirstPath(pathA)
	_, collisions := addSecondPath(pathB, pathACoords)
	shortestCollision := getShortestCollision(collisions)
	return strconv.Itoa(shortestCollision)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
