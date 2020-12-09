package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mapNode map[string]string

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "input", "input file")
	flag.Parse()

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	mapNode = make(map[string]string)

	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), " bags contain ")
		col := spl[0]

		mapNode[col] = scanner.Text()

	}
	fmt.Printf("Sum: %d\n", calc(mapNode["shiny gold"]))

}
func calc(s string) int {
	spl := strings.Split(s, "contain ")
	cont := strings.Split(spl[1], ",")
	sum := 0
	for _, i := range cont {
		if strings.Contains(i, "no other bags") {
			return 0
		}
		spl2 := strings.Split(strings.Trim(i, " ."), " ")
		newBag := (spl2[1] + " " + spl2[2])

		num, _ := strconv.Atoi(spl2[0])

		sum = sum + num + num*calc(mapNode[newBag])

	}
	return sum
}
