package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var max int64
	for scanner.Scan() {
		l := scanner.Text()
		row := l[:7]
		col := l[7:]

		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")

		r, _ := strconv.ParseInt(row, 2, 64)

		col = strings.ReplaceAll(col, "L", "0")
		col = strings.ReplaceAll(col, "R", "1")

		c, _ := strconv.ParseInt(col, 2, 64)

		fmt.Printf("r: %s %d c: %d\n", row, r, c)

		if (r*8 + c) > max {
			max = r*8 + c
		}
	}
	fmt.Println(max)
}
