package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	cnt := 0
	lc := 1
	m := make(map[rune]int)
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			fmt.Println(m)
			fmt.Println(lc)
			for i, _ := range m {
				if m[i] == lc-1 {
					cnt++
				}
			}

			m = make(map[rune]int)
			lc = 0
		}
		for _, r := range l {
			m[r]++
		}
		lc++
	}
	fmt.Println(cnt)
}
