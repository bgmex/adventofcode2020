package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
	var ids []int
	for scanner.Scan() {
		l := scanner.Text()
		res := calc(l)
		ids = append(ids, int(res))

		if res > max {
			max = res
		}
	}

	sort.Ints(ids)
	var cur int
	for j, i := range ids {
		if i != int(1+cur) && j != 0 {
			fmt.Println(i - 1)
		}
		cur = i
	}
}

func calc(l string) int64 {

	row := l[:7]
	col := l[7:]

	row = strings.ReplaceAll(row, "F", "0")
	row = strings.ReplaceAll(row, "B", "1")

	r, _ := strconv.ParseInt(row, 2, 64)

	col = strings.ReplaceAll(col, "L", "0")
	col = strings.ReplaceAll(col, "R", "1")

	c, _ := strconv.ParseInt(col, 2, 64)

	return int64(r*8 + c)
}
