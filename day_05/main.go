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
	memory       []int
	running      bool
	instructions map[int]func(*computer, []int, []int)
	operands     map[int]int
	pcStep       map[int]int
	inputQueue   []int
	outputBuffer []int
}

func newComputer(input []int) computer {
	cpu := computer{0, make([]int, len(input)), true, make(map[int]func(*computer, []int, []int)), make(map[int]int), make(map[int]int), make([]int, 0), make([]int, 0)}
	copy(cpu.memory, input)
	cpu.addInstruction(1, 3, opAdd, 4)
	cpu.addInstruction(2, 3, opMultiply, 4)
	cpu.addInstruction(3, 1, opInput, 2)
	cpu.addInstruction(4, 1, opOutput, 2)
	cpu.addInstruction(5, 2, opJumpIfTrue, 0)
	cpu.addInstruction(6, 2, opJumpIfFalse, 0)
	cpu.addInstruction(7, 3, opLessThan, 4)
	cpu.addInstruction(8, 3, opEquals, 4)
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
	cpu.instructions[opcode](cpu, cpu.memory[cpu.pc+1:cpu.pc+1+cpu.operands[opcode]], addressMode)
	cpu.pc += cpu.pcStep[opcode]
}

func (cpu *computer) addInstruction(opcode int, operands int, f func(*computer, []int, []int), step int) {
	cpu.instructions[opcode] = f
	cpu.operands[opcode] = operands
	cpu.pcStep[opcode] = step
}

func (cpu *computer) read(address int, mode int) int {
	switch mode {
	case 0:
		return cpu.memory[address]
	case 1:
		return address
	}
	return -1
}

func (cpu *computer) write(address int, value int) {
	cpu.memory[address] = value
}

func opAdd(cpu *computer, operands []int, addressMode []int) {
	cpu.write(operands[2], cpu.read(operands[0], addressMode[0])+cpu.read(operands[1], addressMode[1]))
}

func opMultiply(cpu *computer, operands []int, addressMode []int) {
	cpu.write(operands[2], cpu.read(operands[0], addressMode[0])*cpu.read(operands[1], addressMode[1]))
}

func opHalt(cpu *computer, operands []int, addressMode []int) {
	cpu.running = false
}

func opInput(cpu *computer, operands []int, addressMode []int) {
	cpu.write(operands[0], cpu.inputQueue[0])
	cpu.inputQueue = cpu.inputQueue[1:]
}

func opOutput(cpu *computer, operands []int, addressMode []int) {
	value := cpu.read(operands[0], addressMode[0])
	cpu.outputBuffer = append(cpu.outputBuffer, value)
}

func opJumpIfTrue(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0]) != 0 {
		cpu.pc = cpu.read(operands[1], addressMode[1])
	} else {
		cpu.pc += 3
	}
}

func opJumpIfFalse(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0]) == 0 {
		cpu.pc = cpu.read(operands[1], addressMode[1])
	} else {
		cpu.pc += 3
	}
}

func opLessThan(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0]) < cpu.read(operands[1], addressMode[1]) {
		cpu.write(operands[2], 1)
	} else {
		cpu.write(operands[2], 0)
	}
}

func opEquals(cpu *computer, operands []int, addressMode []int) {
	if cpu.read(operands[0], addressMode[0]) == cpu.read(operands[1], addressMode[1]) {
		cpu.write(operands[2], 1)
	} else {
		cpu.write(operands[2], 0)
	}
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

func partA(input string) string {
	cpu := newComputer(splitString(getLines(input)[0]))
	cpu.inputQueue = append(cpu.inputQueue, 1)

	for cpu.running {
		cpu.clock()
	}

	return strconv.Itoa(cpu.outputBuffer[len(cpu.outputBuffer)-1])
}

func partB(input string) string {
	cpu := newComputer(splitString(getLines(input)[0]))
	cpu.inputQueue = append(cpu.inputQueue, 5)

	for cpu.running {
		cpu.clock()
	}

	return strconv.Itoa(cpu.outputBuffer[len(cpu.outputBuffer)-1])
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
