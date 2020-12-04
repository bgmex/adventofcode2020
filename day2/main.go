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
	var cnt int

	flag.StringVar(&inputFile, "i", "", "inputfile")
	flag.Parse()

	file, err := os.Open(inputFile)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if check(line) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func check(line string) bool {
	// line eg:
	// 17-19 p: pwpzpfbrcpppjppbmppp

	spl := strings.Split(line, ":")

	rule := spl[0]
	pass := spl[1][1:]
	char := rule[len(rule)-1:]
	min, err := strconv.Atoi(strings.Split(rule, "-")[0])
	if err != nil {
		return false
	}
	max, err := strconv.Atoi(strings.Split(rule, "-")[1][:len(strings.Split(rule, "-")[1])-2])
	if err != nil {
		return false
	}
	cnt := strings.Count(pass, char)

	if cnt >= min && cnt <= max {
		return true
	}
	return false
}
