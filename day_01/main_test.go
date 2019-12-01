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
			want:  "2",
		},
		{
			input: "a_1.txt",
			want:  "2",
		},
		{
			input: "a_2.txt",
			want:  "654",
		},
		{
			input: "a_3.txt",
			want:  "33583",
		},
		{
			input: "real.txt",
			want:  "3373568",
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
			want:  "5057481",
		},
		{
			input: "b_0.txt",
			want:  "2",
		},
		{
			input: "b_1.txt",
			want:  "966",
		},
		{
			input: "b_2.txt",
			want:  "50346",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := partB(test.input)

			assert.Equal(test.want, got)
		})
	}
}
