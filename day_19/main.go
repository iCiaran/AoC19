package main

import (
	"bufio"
	"fmt"
	"log"
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
		} else {
			return cpu.memory[address]
		}
	case 1:
		return address
	case 2:
		if write {
			return cpu.base + address
		} else {
			return cpu.memory[cpu.base+address]
		}
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

func isPulled(cpu *computer, c coord) int {
	cpu.inputQueue = append(cpu.inputQueue, []int{c.x, c.y}...)
	for cpu.running && len(cpu.outputBuffer) < 1 {
		cpu.clock()
	}
	res := cpu.outputBuffer[0]
	cpu.outputBuffer = nil
	return res
}

func findFirstXFit(input string, width int) int {
	split := splitString(getLines(input)[0])
	start := 0
	end := 1000
	guess := 0
	for (end-start)/2 > 0 {
		guess = start + (end-start)/2
		total := 0
		for i := 0; i < guess; i++ {
			cpu := newComputer(split)
			total += isPulled(&cpu, coord{i, guess})
		}
		if total >= width {
			end = guess
		} else {
			start = guess
		}
	}

	return end
}

func findFirstXYFit(input string, startY int) (int, int) {
	split := splitString(getLines(input)[0])
	found := false
	y := startY
	x := 1

	for !found {
		cpu := newComputer(split)
		previous := isPulled(&cpu, coord{0, y})
		cpu = newComputer(split)
		current := isPulled(&cpu, coord{1, y})
		for !(current == 0 && previous == 1) {
			x++
			previous = current
			cpu = newComputer(split)
			current = isPulled(&cpu, coord{x, y})
		}
		x -= 100
		cpu = newComputer(split)
		if isPulled(&cpu, coord{x, y + 99}) == 1 {
			found = true
		} else {
			y++
		}
	}
	return x, y
}

func partA(input string) string {
	count := 0
	split := splitString(getLines(input)[0])
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			cpu := newComputer(split)
			count += isPulled(&cpu, coord{x, y})
		}
	}
	return strconv.Itoa(count)
}

func partB(input string) string {
	startY := findFirstXFit(input, 100)
	x, y := findFirstXYFit(input, startY)
	return strconv.Itoa(x*10000 + y)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
