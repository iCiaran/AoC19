package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
			want:  "A",
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
			want:  "B",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := partB(test.input)

			assert.Equal(test.want, got)
		})
	}
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
