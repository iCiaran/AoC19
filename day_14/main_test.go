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
			want:  "399063",
		},
		{
			input: "a_0.txt",
			want:  "165",
		},
		{
			input: "a_1.txt",
			want:  "13312",
		},
		{
			input: "a_2.txt",
			want:  "180697",
		},
		{
			input: "a_3.txt",
			want:  "2210736",
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
			want:  "4215654",
		},
		{
			input: "a_1.txt",
			want:  "82892753",
		},
		{
			input: "a_2.txt",
			want:  "5586022",
		},
		{
			input: "a_3.txt",
			want:  "460664",
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
