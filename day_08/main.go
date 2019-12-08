package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

func buildLayers(input string) [][]int {
	const pixels int = 25 * 6
	image := make([][]int, len(input)/pixels)
	for i := range image {
		image[i] = make([]int, pixels)
	}

	for i, c := range input {
		image[i/pixels][i%pixels] = int(c) - '0'
	}
	return image
}

func combineLayers(image [][]int) []int {
	const pixels int = 25 * 6
	combined := make([]int, pixels)

	for layer := len(image) - 1; layer >= 0; layer-- {
		for i, p := range image[layer] {
			if p == 0 || p == 1 {
				combined[i] = p
			}
		}
	}

	return combined
}

func count(image [][]int) int {
	min := math.MaxInt32
	result := 0
	for _, layer := range image {
		counts := make([]int, 10)
		for _, pixel := range layer {
			counts[pixel]++
		}
		if counts[0] < min {
			min = counts[0]
			result = counts[1] * counts[2]
		}
	}
	return result
}

func formatLayer(layer []int) string {
	width := 25
	height := 6
	var sb strings.Builder
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch layer[y*width+x] {
			case 0:
				sb.WriteString(".")
			case 1:
				sb.WriteString("#")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func partA(input string) string {
	return strconv.Itoa(count(buildLayers(getLines(input)[0])))
}

func partB(input string) string {
	output := combineLayers(buildLayers(getLines(input)[0]))
	return formatLayer(output)
}

func main() {
	input := "real.txt"
	fmt.Println(partA(input))
	fmt.Println(partB(input))
}
