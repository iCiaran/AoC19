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
	instructions map[int]func(*computer, []int)
	operands     map[int]int
	pcStep       map[int]int
}

func newComputer(input []int) computer {
	cpu := computer{0, make([]int, len(input)), true, make(map[int]func(*computer, []int)), make(map[int]int), make(map[int]int)}
	copy(cpu.memory, input)
	cpu.addInstruction(1, 3, add, 4)
	cpu.addInstruction(2, 3, multiply, 4)
	cpu.addInstruction(99, 0, halt, 1)
	return cpu
}

func (cpu *computer) clock() {
	opcode := cpu.memory[cpu.pc]
	cpu.instructions[opcode](cpu, cpu.memory[cpu.pc+1:cpu.pc+1+cpu.operands[opcode]])
	cpu.pc += cpu.pcStep[opcode]
}

func (cpu *computer) addInstruction(opcode int, operands int, f func(*computer, []int), step int) {
	cpu.instructions[opcode] = f
	cpu.operands[opcode] = operands
	cpu.pcStep[opcode] = step
}

func add(cpu *computer, operands []int) {
	cpu.memory[operands[2]] = cpu.memory[operands[0]] + cpu.memory[operands[1]]
}

func multiply(cpu *computer, operands []int) {
	cpu.memory[operands[2]] = cpu.memory[operands[0]] * cpu.memory[operands[1]]
}

func halt(cpu *computer, operands []int) {
	cpu.running = false
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
	cpu.memory[1] = 12
	cpu.memory[2] = 2

	for cpu.running {
		cpu.clock()
	}

	return strconv.Itoa(cpu.memory[0])
}

func partB(input string) string {
	target := 19690720
	initial := splitString(getLines(input)[0])

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			cpu := newComputer(initial)
			cpu.memory[1] = noun
			cpu.memory[2] = verb
			for cpu.running {
				cpu.clock()
			}
			if cpu.memory[0] == target {
				return strconv.Itoa(100*noun + verb)
			}
		}
	}

	return "Not found"
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
