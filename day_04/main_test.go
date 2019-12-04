package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleCheck(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 111111,
			want:  true,
		},
		{
			input: 223450,
			want:  false,
		},
		{
			input: 123789,
			want:  false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := check(test.input)

			assert.Equal(test.want, got)
		})
	}
}

func TestRuleCheckNoGroups(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 112233,
			want:  true,
		},
		{
			input: 123444,
			want:  false,
		},
		{
			input: 111122,
			want:  true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := checkNoGroups(test.input)

			assert.Equal(test.want, got)
		})
	}
}
func TestPartA(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "real.txt",
			want:  "579",
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
			want:  "358",
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
