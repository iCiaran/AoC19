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

type computer struct {
	pc           int
	memory       map[int]int
	running      bool
	instructions map[int]func(*computer, []int, []int)
	operands     map[int]int
	pcStep       map[int]int
	inputQueue   []int
	outputBuffer []int
	base         int
	paused       bool
}

type coord struct {
	x int
	y int
}

type search struct {
	c coord
	d int
}

func newComputer(input []int) computer {
	cpu := computer{0, make(map[int]int, len(input)), true, make(map[int]func(*computer, []int, []int)), make(map[int]int), make(map[int]int), make([]int, 0), make([]int, 0), 0, false}
	for key, val := range input {
		cpu.memory[key] = val
	}
	cpu.addInstruction(1, 3, opAdd, 4)
	cpu.addInstruction(2, 3, opMultiply, 4)
	cpu.addInstruction(3, 1, opInput, 2)
	cpu.addInstruction(4, 1, opOutput, 2)
	cpu.addInstruction(5, 2, opJumpIfTrue, 0)
	cpu.addInstruction(6, 2, opJumpIfFalse, 0)
	cpu.addInstruction(7, 3, opLessThan, 4)
	cpu.addInstruction(8, 3, opEquals, 4)
	cpu.addInstruction(9, 1, opBase, 2)
	cpu.addInstruction(99, 0, opHalt, 1)
	return cpu
}

func (cpu *computer) clock() {
	opcode := cpu.memory[cpu.pc] % 100
	var addressMode []int
	div := 100
	for i := 0; i < cpu.operands[opcode]; i++ {
		addressMode = append(addressMode, (cpu.memory[cpu.pc]/div)%10)
		div *= 10
	}
	cpu.instructions[opcode](cpu, cpu.getOperands(cpu.pc, cpu.operands[opcode]), addressMode)
	cpu.pc += cpu.pcStep[opcode]
}

func (cpu *computer) addInstruction(opcode int, operands int, f func(*computer, []int, []int), step int) {
	cpu.instructions[opcode] = f
	cpu.operands[opcode] = operands
	cpu.pcStep[opcode] = step
}

func (cpu *computer) getOperands(address int, length int) []int {
	operands := make([]int, length)
	for i := 0; i < length; i++ {
		operands[i] = cpu.memory[cpu.pc+i+1]
	}
	return operands
}

func (cpu *computer) read(address int, mode int, write bool) int {
	switch mode {
	case 0:
		if write {
			return address
		}
		return cpu.memory[address]
	case 1:
		return address
	case 2:
		if write {
			return cpu.base + address
		}
		return cpu.memory[cpu.base+address]
	default:
		log.Fatal("Wrong mode: ", mode)
		os.Exit(1)
	}
	return -1
}

func (cpu *computer) write(address int, value int) {
	cpu.memory[address] = value
}

func opAdd(cpu *computer, operands []int, addressMode []int) {
	cpu.write(cpu.read(operands[2], addressMode[2], true), cpu.read(operands[0], addressMode[0], false)+cpu.read(operands[1], addressMode[1], false))
}

func opMultiply(cpu *computer, operands []int, addressMode []int) {
	cpu.write(cpu.read(operands[2], addressMode[2], true), cpu.read(operands[0], addressMode[0], false)*cpu.read(operands[1], addressMode[1], false))
}

func opHalt(cpu *computer, operands []int, addressMode []int) {
	cpu.running = false
}

func opInput(cpu *computer, operands []int, addressMode []int) {
	if len(cpu.inputQueue) > 0 {
		cpu.write(cpu.read(operands[0], addressMode[0], true), cpu.inputQueue[0])
		cpu.inputQueue = cpu.inputQueue[1:]
	} else {
		cpu.paused = true
		cpu.pc -= 2
	}
}

func opOutput(cpu *computer, operands []int, addressMode []int) {
	value := cpu.read(operands[0], addressMode[0], false)
	cpu.outputBuffer = append(cpu.outputBuffer, value)
}

func opJumpIfTrue(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0], false) != 0 {
		cpu.pc = cpu.read(operands[1], addressMode[1], false)
	} else {
		cpu.pc += 3
	}
}

func opJumpIfFalse(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0], false) == 0 {
		cpu.pc = cpu.read(operands[1], addressMode[1], false)
	} else {
		cpu.pc += 3
	}
}

func opLessThan(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0], false) < cpu.read(operands[1], addressMode[1], false) {
		cpu.write(cpu.read(operands[2], addressMode[2], true), 1)
	} else {
		cpu.write(cpu.read(operands[2], addressMode[2], true), 0)
	}
}

func opEquals(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0], false) == cpu.read(operands[1], addressMode[1], false) {
		cpu.write(cpu.read(operands[2], addressMode[2], true), 1)
	} else {
		cpu.write(cpu.read(operands[2], addressMode[2], true), 0)
	}
}

func opBase(cpu *computer, operands []int, addressMode []int) {
	cpu.base += cpu.read(operands[0], addressMode[0], false)
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

func splitString(input string) []int {
	tmp := strings.Split(input, ",")
	var values []int

	for _, i := range tmp {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		values = append(values, num)
	}
	return values
}

func getMove(cpu *computer, r, i int) int {
	var command int

	if r == 0 && i == 1 {
		command = 1
	} else if r == 0 && i == -1 {
		command = 2
	} else if r == 1 && i == 0 {
		command = 3
	} else if r == -1 && i == 0 {
		command = 4
	}

	cpu.inputQueue = append(cpu.inputQueue, command)
	for len(cpu.outputBuffer) < 1 {
		cpu.clock()
	}

	result := cpu.outputBuffer[0]
	cpu.outputBuffer = cpu.outputBuffer[:0]

	return result
}

func getBounds(system map[coord]rune) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for k := range system {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	return minX, minY, maxX, maxY
}

func getOutputString(system map[coord]rune) string {
	minX, minY, maxX, maxY := getBounds(system)
	var sb strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c, i := system[coord{x, y}]
			if x == 0 && y == 0 {
				c = 'o'
			}
			if i {
				sb.WriteRune(c)
			} else {
				sb.WriteRune(' ')
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func buildMap(input string) (map[coord]rune, coord) {
	location := coord{0, 0}
	origin := coord{0, 0}
	right := complex(0, 1)
	left := complex(0, -1)
	target := coord{0, 0}
	direction := complex(0, 1)
	count := 0
	system := make(map[coord]rune)
	system[location] = '.'
	cpu := newComputer(splitString(getLines(input)[0]))

	for count < 4 {
		next := coord{location.x + int(real(direction)), location.y + int(imag(direction))}
		if next == origin {
			count++
		}
		res := getMove(&cpu, int(real(direction)), int(imag(direction)))
		if res == 0 {
			system[next] = '#'
			direction *= right
		} else if res == 1 {
			system[next] = '.'
			direction *= left
			location = next
		} else if res == 2 {
			system[next] = 'x'
			direction *= left
			location = next
			target = next
		}
	}

	return system, target
}

func getNeighbours(maze map[coord]rune, current coord) []coord {
	n := make([]coord, 0)
	list := [4]coord{coord{current.x + 1, current.y}, coord{current.x - 1, current.y}, coord{current.x, current.y + 1}, coord{current.x, current.y - 1}}
	for _, c := range list {
		if maze[c] == '.' || maze[c] == 'x' {
			n = append(n, c)
		}
	}
	return n
}

func getDistance(maze map[coord]rune, target rune) int {
	queue := make([]search, 0)
	queue = append(queue, search{coord{0, 0}, 0})
	visited := make(map[coord]bool, 0)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current.c] = true
		for _, c := range getNeighbours(maze, current.c) {
			if maze[c] == 'x' {
				return current.d + 1
			} else if !visited[c] {
				queue = append(queue, search{c, current.d + 1})
			}
		}
	}

	return -1
}

func getTime(maze map[coord]rune, oxygen coord) int {
	queue := make([]search, 0)
	queue = append(queue, search{oxygen, 0})
	last := 0
	visited := make(map[coord]bool, 0)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current.c] = true
		last = current.d
		for _, c := range getNeighbours(maze, current.c) {
			if !visited[c] {
				queue = append(queue, search{c, current.d + 1})
			}
		}
	}

	return last
}

func partA(input string) string {
	maze, _ := buildMap(input)
	distance := getDistance(maze, 'x')
	return strconv.Itoa(distance)
}

func partB(input string) string {
	maze, oxygen := buildMap(input)
	time := getTime(maze, oxygen)
	return strconv.Itoa(time)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
