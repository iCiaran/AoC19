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
			want:  "63794407",
		},
		{
			input: "a_0.txt",
			want:  "23845678",
		},
		{
			input: "a_1.txt",
			want:  "24176176",
		},
		{
			input: "a_2.txt",
			want:  "73745418",
		},
		{
			input: "a_3.txt",
			want:  "52432133",
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
			want:  "77247538",
		},
		{
			input: "b_0.txt",
			want:  "84462026",
		},
		{
			input: "b_1.txt",
			want:  "78725270",
		},
		{
			input: "b_2.txt",
			want:  "53553731",
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
