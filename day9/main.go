package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

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
		i++
		if i <= 25 {
			m[i%25] = n
			continue
		}
		for oi := 0; oi < 25; oi++ {
			for ii := 0; ii < 25; ii++ {
				if ii == oi {
					continue
				}
				if m[ii]+m[oi] == n {
					goto Cont
				}
			}
		}
		fmt.Println(n)
	Cont:
		m[i%25] = n
	}
}
