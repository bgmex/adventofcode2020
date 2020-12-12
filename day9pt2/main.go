package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// SOLPT1 is the solution from part one
var SOLPT1 int64 = 50047984

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "input", "input file")
	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var i int
	m := make(map[int]int64)
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		m[i] = n
		i++
	}
	calc(m, SOLPT1)
}
func calc(m map[int]int64, n int64) {

	for i := 0; i < len(m); i++ {
		acc := m[i]
		for j := i + 1; j < len(m); j++ {
			acc += m[j]
			if acc == n {
				fmt.Printf("i: %d , j: %d\n", i, j)
				v := make([]int64, j-i+1, j-i+1)
				for I := i; I <= j; I++ {
					fmt.Printf("m[%d]: %d v[%d]\n", I, m[I], I-i)
					v[I-i] = m[I]
				}
				sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
				fmt.Println("Solution:")
				fmt.Println(v[0] + v[len(v)-1])
			}
			if acc > n {
				continue
			}
		}

	}
	os.Exit(0)
}
