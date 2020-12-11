package main

import (
	"bufio"
	"flag"
	"fmt"

	"os"
	"strconv"
	"strings"
)

type instr struct {
	op      string
	val     int
	visited bool
}

var m = make(map[int]instr)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "input", "input file")
	flag.Parse()

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var line, maxnop, maxjmp int
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(spl[1])
		in := instr{spl[0], val, false}
		m[line] = in
		line++
		if spl[0] == "nop" {
			maxnop++
		}
		if spl[0] == "jmp" {
			maxjmp++
		}

	}
	fmt.Println(run())
}

func run() int {
	var line, acc int
	for {
		if line == len(m) {
			fmt.Printf("accumulator: %d\n", acc)
			os.Exit(0)
		}
		m[line] = instr{m[line].op, m[line].val, true}
		switch m[line].op {
		case "acc":
			acc += m[line].val
			line++
		case "jmp":
			line += m[line].val
			if m[line].visited {
				return acc
			}
		case "nop":
			line++
		}

	}
}
