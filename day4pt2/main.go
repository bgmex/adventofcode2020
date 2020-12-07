package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	Name string
	Test func(string) bool
}

var fields []Field = []Field{
	{"byr", checkBYR},
	{"iyr", checkIYR},
	{"eyr", checkEYR},
	{"hgt", checkHGT},
	{"hcl", checkHCL},
	{"ecl", checkECL},
	{"pid", checkPID},
}

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "inputfile", "input file")
	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var cnt int
	var s string

	// empty line in input file
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			cnt += analyse(s + " ")
			s = ""
		}
		s = s + " " + scanner.Text()
	}
	fmt.Println(cnt)
}

func analyse(s string) int {

	for _, f := range fields {
		idx := strings.Index(s, f.Name)
		if idx == -1 {
			return 0
		}
		idx2 := strings.Index(s[idx:], " ")
		if idx2 == -1 {
			return 0
		}
		spl := strings.Split(s[idx:idx+idx2], ":")
		if !f.Test(spl[1]) {
			return 0
		}
	}
	fmt.Println(s)
	return 1
}
func checkYR(s string, min int, max int) bool {
	yr, err := strconv.Atoi(s)

	if err != nil || yr < min || yr > max {
		return false
	}

	return true
}
func checkBYR(s string) bool {
	return checkYR(s, 1920, 2002)
}
func checkIYR(s string) bool {
	return checkYR(s, 2010, 2020)
}
func checkEYR(s string) bool {
	return checkYR(s, 2020, 2030)
}
func checkHGT(s string) bool {
	idxIN := strings.Index(s, "in")
	idxCM := strings.Index(s, "cm")

	hgt, err := strconv.Atoi(s[:len(s)-2])
	if err != nil || idxCM*idxIN > 0 {
		return false
	}

	if idxCM != -1 && (hgt < 150 || hgt > 193) {
		return false
	}
	if idxIN != -1 && (hgt < 59 || hgt > 76) {
		return false
	}
	return true
}
func checkHCL(s string) bool {
	if s[0] != '#' || len(s) != 7 || s[2] == 'x' || s[2] == 'X' {
		return false
	}
	_, err := strconv.ParseUint(s[1:], 16, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func checkECL(s string) bool {
	cols := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, col := range cols {
		if col == s {
			return true
		}
	}
	return false
}
func checkPID(s string) bool {
	if len(s) != 9 {
		return false
	}

	for _, i := range s {
		if i-'0' < 0 || i-'0' > 9 {
			return false
		}
	}
	return true
}
