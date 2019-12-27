package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type packet struct {
	x int
	y int
}

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
		cpu.paused = false
	} else { //Modified for today to not block and instead give -1
		cpu.write(cpu.read(operands[0], addressMode[0], true), -1)
		cpu.paused = true
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

func partA(input string) string {
	split := splitString(getLines(input)[0])
	found := false
	result := 0
	computers := make([]computer, 50)
	for i := 0; i < 50; i++ {
		computers[i] = newComputer(split)
		computers[i].inputQueue = append(computers[i].inputQueue, i)
	}

	for !found {
		for i := 0; i < 50; i++ {
			computers[i].clock()
			if len(computers[i].outputBuffer) >= 3 {
				if computers[i].outputBuffer[0] == 255 {
					found = true
					result = computers[i].outputBuffer[2]
				} else {
					computers[computers[i].outputBuffer[0]].inputQueue = append(computers[computers[i].outputBuffer[0]].inputQueue, computers[i].outputBuffer[1:3]...)
				}
				computers[i].outputBuffer = computers[i].outputBuffer[3:]
			}
		}
	}
	return strconv.Itoa(result)
}

func partB(input string) string {
	split := splitString(getLines(input)[0])
	found := false
	result := 0
	computers := make([]computer, 50)

	NATPacket := packet{0, 0}
	NATLastY := -1

	for i := 0; i < 50; i++ {
		computers[i] = newComputer(split)
		computers[i].inputQueue = append(computers[i].inputQueue, i)
	}

	for !found {
		for i := 0; i < 50; i++ {
			computers[i].clock()
			if len(computers[i].outputBuffer) >= 3 {
				if computers[i].outputBuffer[0] == 255 {
					NATPacket.x = computers[i].outputBuffer[1]
					NATPacket.y = computers[i].outputBuffer[2]
				} else {
					computers[computers[i].outputBuffer[0]].inputQueue = append(computers[computers[i].outputBuffer[0]].inputQueue, computers[i].outputBuffer[1:3]...)
				}
				computers[i].outputBuffer = computers[i].outputBuffer[3:]
			}
		}

		count := 0
		for i := 0; i < 50; i++ {
			if computers[i].paused && len(computers[i].inputQueue) == 0 {
				count++
			}
		}

		if count == 50 {
			if NATPacket.y == NATLastY {
				found = true
				result = NATLastY
			}
			computers[0].inputQueue = append(computers[0].inputQueue, []int{NATPacket.x, NATPacket.y}...)
			NATLastY = NATPacket.y
		}
	}

	return strconv.Itoa(result)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
