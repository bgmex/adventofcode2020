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
	m := make(map[rune]int)
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			cnt += len(m)
			m = make(map[rune]int)
		}
		for _, r := range l {
			m[r]++
		}
	}
	fmt.Println(cnt)
}
