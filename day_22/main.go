package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

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

func getCoefficients(line string) (a, b int) {
	split := strings.Split(line, " ")
	if split[0] == "cut" {
		a = 1
		b,_ = strconv.Atoi(split[1])
		b *= -1
	} else if split[1] == "into" {
		a = -1
		b = -1
	} else {
		a,_ = strconv.Atoi(split[3])
		b = 0
	}
	return
}

func compose(m int, input string) (int,int) {
	lines := getLines(input)
	a, b := getCoefficients(lines[0])

	for i := 1; i < len(lines); i++ {
		c, d := getCoefficients(lines[i])
		a = mod((a * c), m)
		b = mod((b * c + d),m)
	}
	return a,b
}

func getPosition(m, x int, input string) int {
	a,b := compose(m, input)
	return mod((a * x + b),m)
}

func mod(a, b int) int {
	res := a % b
	for res < 0 {
		res += b
	}
	return res
}

func modInverse(a,b int) int {
	t := 0
	newt := 1
	r := b
	newr := a

	for newr != 0 {
		q := r / newr
		tempt := t
		t = newt
		newt = tempt - q * newt

		tempr := r
		r = newr
		newr = tempr - q * newr
	}

	if r > 1 {
		log.Fatal("Not invertible: ", a ,b)
		os.Exit(1)
	}

	if t < 0 {
		t += b
	}

	return t
}

func expBySquaring(x, n, m int) int {
	x2 := mod(x*x, m)
	if n < 0 {
		return expBySquaring(1/x, -n,m)
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n % 2 == 0 {
		return expBySquaring(x2, n/2, m)
	} else {
		return expBySquaring(x2, (n-1)/2, m)
	}
}

func partA(input string) string {
	return strconv.Itoa(getPosition(10007, 2019, input))
}

func partB(input string) string {
	m := 119315717514047
	k := 101741582076661
	x := 2020
	a, b := compose(m, input)

	ak := mod(expBySquaring(a,k,m), m)
	inv := modInverse(b *  (1-ak), 1-a)

	res := mod(modInverse(x - inv, ak), m)
	return strconv.Itoa(res)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
