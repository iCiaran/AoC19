package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type coord struct {
	x float64
	y float64
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

func getAsteroids(input []string) []coord {
	asteroids := make([]coord, 0)
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				a := coord{float64(x), float64(len(input) - y - 1)}
				asteroids = append(asteroids, a)
			}
		}
	}

	return asteroids
}

func getAngles(asteroids []coord, from coord) map[float64]map[coord]bool {
	angles := make(map[float64]map[coord]bool, 0)
	for _, to := range asteroids {
		if from != to {
			angle := math.Atan2(to.y-from.y, to.x-from.x)
			if _, ok := angles[angle]; !ok {
				angles[angle] = make(map[coord]bool, 0)
			}
			angles[angle][to] = true
		}
	}
	return angles
}

func maxVisible(asteroids []coord) (int, coord) {
	max := 0
	maxCoord := coord{-1.0, -1.0}
	for _, a := range asteroids {
		visible := len(getAngles(asteroids, a))
		if visible > max {
			max = visible
			maxCoord = a
		}
	}
	return max, maxCoord
}

func getOrder(angles map[float64]map[coord]bool) []float64 {
	order := make([]float64, len(angles))
	i := 0
	for key := range angles {
		order[i] = key
		i++
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(order)))
	cutoff := 0
	for order[cutoff] > (math.Pi / 2.0) {
		cutoff++
	}

	order = append(order[cutoff:], order[:cutoff]...)
	return order
}

func getClosest(asteroids []coord) coord {
	closest := math.MaxFloat64
	result := coord{0.0, 0.0}
	for _, a := range asteroids {
		distance2 := a.x*a.x + a.y + a.y
		if distance2 < closest {
			closest = distance2
			result = a
		}
	}
	return result
}

func nthDestroyed(order []float64, angles map[float64]map[coord]bool, n int) coord {
	i := 0
	destroyed := coord{-1, -1}
	for i < n {
		angle := order[i%n]
		toCheck := make([]coord, 0)
		for k, v := range angles[angle] {
			if v {
				toCheck = append(toCheck, k)
			}
		}

		if len(toCheck) > 0 {
			destroyed = getClosest(toCheck)
			angles[angle][destroyed] = false
			i++
		}
	}
	return destroyed
}

func partA(input string) string {
	asteroids := getAsteroids(getLines(input))
	num, _ := maxVisible(asteroids)
	return strconv.Itoa(num)
}

func partB(input string) string {
	lines := getLines(input)
	asteroids := getAsteroids(lines)
	_, a := maxVisible(asteroids)
	angles := getAngles(asteroids, a)
	order := getOrder(angles)
	last := nthDestroyed(order, angles, 200)
	x := int(last.x)
	y := len(lines) - 1 - int(last.y)
	return strconv.Itoa(x*100 + y)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
