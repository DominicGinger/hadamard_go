package main

import (
	"fmt"
	"math/rand"
	"time"
)

var h1 = make([][]int, 2)

func setup() {
	h1[0] = []int{1, 1}
	h1[1] = []int{1, -1}
}

func pow(base int, p int) int {
	n := base
	for i := 1; i < p; i++ {
		n *= base
	}
	return n
}

func buildMatrix(num int) [][]int {
	base := h1
	for k := 2; k <= num; k++ {
		size := pow(2, k)
		half := size / 2
		h := make([][]int, size)
		for i, _ := range h {
			h[i] = make([]int, size)
			for j, _ := range h[i] {
				x, y := i, j
				if i >= half {
					x -= half
				}
				if j >= half {
					y -= half
				}

				if i < half || j < half {
					h[i][j] = base[x][y]
				} else {
					h[i][j] = -base[x][y]
				}
			}
		}
		base = h
	}
	return base
}

func printM(h [][]int) {
	for _, v := range h {
		for _, x := range v {
			if x == 1 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("# ")
			}
		}
		fmt.Println()
	}
}

func decode(matrix [][]int, encoded [][]int) string {
	decoded := ""
	for _, message := range encoded {
		match := -1
		biggestCount := 0
		for i, v := range matrix {
			count := 0
			for j, _ := range v {
				if j < len(message) && v[j] == message[j] {
					count++
				}
			}
			if count > biggestCount {
				biggestCount = count
				match = i
			}
		}
		decoded += string(match)
	}
	return decoded
}

func encode(matrix [][]int, message string) [][]int {
	encoded := [][]int{}
	for _, r := range message {
		e := make([]int, len(matrix[r]))
		copy(e, matrix[r])
		encoded = append(encoded, e)
	}
	return encoded
}

func inArray(arr []int, a int) bool {
	for _, v := range arr {
		if v == a {
			return true
		}
	}
	return false
}

func scramble(encoded [][]int, noBits int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range encoded {
		scrambled := []int{}
		for j := 0; j < noBits; j++ {
			nextRand := rand.Intn(len(encoded[i]))
			for inArray(scrambled, nextRand) {
				nextRand = rand.Intn(len(encoded[i]))
			}
			encoded[i][nextRand] *= -1
		}
	}
}

func main() {
	setup()
	h := buildMatrix(8)
	inp := "The quick brown fox jumps over the lazy dog"
	encoded := encode(h, inp)
	scramble(encoded, 178)
	fmt.Println(decode(h, encoded))
}
