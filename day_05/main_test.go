package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestPartA(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "real.txt",
			want:  "7839346",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := partA(test.input)

			assert.Equal(test.want, got)
		})
	}
}

func TestPartB(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "real.txt",
			want:  "447803",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := partB(test.input)

			assert.Equal(test.want, got)
		})
	}
}

func TestJumps(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		inputFile string
		inputNum  int
		want      string
	}{
		{
			"jumps_0.txt",
			0,
			"0",
		},
		{
			"jumps_0.txt",
			10,
			"1",
		},
		{
			"jumps_1.txt",
			0,
			"0",
		},
		{
			"jumps_1.txt",
			10,
			"1",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			cpu := newComputer(splitString(getLines(test.inputFile)[0]))
			cpu.inputQueue = append(cpu.inputQueue, test.inputNum)

			for cpu.running {
				cpu.clock()
			}

			assert.Equal(test.want, strconv.Itoa(cpu.outputBuffer[len(cpu.outputBuffer)-1]))
		})
	}
}

func TestComparisons(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		inputFile string
		inputNum  int
		want      string
	}{
		{
			"comps_0.txt",
			7,
			"999",
		},
		{
			"comps_0.txt",
			8,
			"1000",
		},
		{
			"comps_0.txt",
			9,
			"1001",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			cpu := newComputer(splitString(getLines(test.inputFile)[0]))
			cpu.inputQueue = append(cpu.inputQueue, test.inputNum)

			for cpu.running {
				cpu.clock()
			}

			assert.Equal(test.want, strconv.Itoa(cpu.outputBuffer[len(cpu.outputBuffer)-1]))
		})
	}
}

func TestBackCompatibility(t *testing.T) {
	t.Run("test_0", func(t *testing.T) {
		cpu := newComputer(splitString(getLines("day_02_real.txt")[0]))
		cpu.write(1, 12)
		cpu.write(2, 2)

		for cpu.running {
			cpu.clock()
		}

		assert.Equal(t, "3931283", strconv.Itoa(cpu.read(0, 0)))
	})
}

func BenchmarkPartA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		partA("real.txt")
	}
}

func BenchmarkPartB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		partB("real.txt")
	}
}
