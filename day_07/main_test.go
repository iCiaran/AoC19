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
			want:  "255590",
		},
		{
			input: "a_0.txt",
			want:  "43210",
		},
		{
			input: "a_1.txt",
			want:  "54321",
		},
		{
			input: "a_2.txt",
			want:  "65210",
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
			want:  "58285150",
		},
		{
			input: "b_0.txt",
			want:  "139629729",
		},
		{
			input: "b_1.txt",
			want:  "18216",
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
