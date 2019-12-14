package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	a int
	b int
}

type state struct {
	p1, p2, p3, p4, v1, v2, v3, v4 int
}

type moon struct {
	position []int
	velocity []int
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

func calculateVelocity(moons []moon, pairs []pair, axis int) {
	for _, p := range pairs {
		if moons[p.a].position[axis] > moons[p.b].position[axis] {
			moons[p.a].velocity[axis] -= 1
			moons[p.b].velocity[axis] += 1
		} else if moons[p.a].position[axis] < moons[p.b].position[axis] {
			moons[p.a].velocity[axis] += 1
			moons[p.b].velocity[axis] -= 1
		}
	}
}

func applyVelocity(moons []moon, axis int) {
	for i, _ := range moons {
		moons[i].position[axis] += moons[i].velocity[axis]
	}
}

func printMoons(moons []moon) {
	for i, _ := range moons {
		fmt.Printf("pos=<x= %3d, y= %3d, z= %3d>, vel=<x= %3d, y= %3d, z= %3d>\n", moons[i].position[0], moons[i].position[1], moons[i].position[2], moons[i].velocity[0], moons[i].velocity[1], moons[i].velocity[2])
	}
	fmt.Println()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getEnergy(m moon) int {
	return (abs(m.position[0]) + abs(m.position[1]) + abs(m.position[2])) * (abs(m.velocity[0]) + abs(m.velocity[1]) + abs(m.velocity[2]))
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func setup(input string) ([]moon, []pair) {
	lines := getLines(input)

	moons := make([]moon, 0)
	for _, line := range lines {
		s := strings.Split(line[1:len(line)-1], ",")
		x, _ := strconv.Atoi(strings.Split(s[0], "=")[1])
		y, _ := strconv.Atoi(strings.Split(s[1], "=")[1])
		z, _ := strconv.Atoi(strings.Split(s[2], "=")[1])
		m := moon{make([]int, 0), make([]int, 0)}
		m.position = append(m.position, []int{x, y, z}...)
		m.velocity = append(m.velocity, []int{0, 0, 0}...)
		moons = append(moons, m)
	}

	pairs := make([]pair, 0)
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			pairs = append(pairs, pair{i, j})
		}
	}

	return moons, pairs
}

func partA(input string) string {

	moons, pairs := setup(input)
	for i := 0; i < 1000; i++ {
		for axis := 0; axis < 3; axis++ {
			calculateVelocity(moons, pairs, axis)
			applyVelocity(moons, axis)
		}
	}

	total := 0
	for _, m := range moons {
		total += getEnergy(m)
	}

	return strconv.Itoa(total)
}

func partB(input string) string {
	moons, pairs := setup(input)

	states := make([]map[state]bool, 3)
	states[0] = make(map[state]bool)
	states[1] = make(map[state]bool)
	states[2] = make(map[state]bool)

	count := make([]int, 3)

	for i := 0; i < 3; i++ {
		current := state{moons[0].position[i], moons[1].position[i], moons[2].position[i], moons[3].position[i], moons[0].velocity[i], moons[1].velocity[i], moons[2].velocity[i], moons[3].velocity[i]}
		for !states[i][current] {
			calculateVelocity(moons, pairs, i)
			applyVelocity(moons, i)

			states[i][current] = true
			count[i]++
			current = state{moons[0].position[i], moons[1].position[i], moons[2].position[i], moons[3].position[i], moons[0].velocity[i], moons[1].velocity[i], moons[2].velocity[i], moons[3].velocity[i]}
		}
	}

	return strconv.Itoa(lcm(count[0], lcm(count[1], count[2])))
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
