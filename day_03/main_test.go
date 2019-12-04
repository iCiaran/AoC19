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
			input: "a_0.txt",
			want:  "6",
		},
		{
			input: "a_1.txt",
			want:  "159",
		},
		{
			input: "a_2.txt",
			want:  "135",
		},
		{
			input: "real.txt",
			want:  "1519",
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
			input: "a_0.txt",
			want:  "30",
		},
		{
			input: "a_1.txt",
			want:  "610",
		},
		{
			input: "a_2.txt",
			want:  "410",
		},
		{
			input: "real.txt",
			want:  "14358",
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

func BenchmarkTogether(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bothTogether("real.txt")
	}
}
