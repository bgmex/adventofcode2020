package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var inputFile string
	var right, down int

	flag.StringVar(&inputFile, "i", "", "input file")
	flag.IntVar(&right, "r", 0, "squares to the right")
	flag.IntVar(&down, "d", 0, "squares down")

	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	if down == 0 {
		fmt.Println("Unsolvable")
		return
	}

	fmt.Println(solve(file, right, down))
}

func solve(i io.Reader, r int, down int) int {
	scanner := bufio.NewScanner(i)

	var cnt, row, right int
	right = r
	// scann first line in order to find width
	scanner.Scan()
	width := len(scanner.Text())

	if scanner.Text()[0] == '#' {
		cnt++
	}
	row++

	for scanner.Scan() {
		if row%down == 0 {
			if scanner.Text()[right%width] == '#' {
				cnt++
			}
			right = (right + r) % width
		}
		row++
	}
	return cnt
}
