package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	mapNode := make(map[string]string)

	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), " bags contain ")
		col := spl[0]

		mapNode[col] = spl[1]

	}

	var last int

	finder := make(map[string]bool)
	finder["shiny gold"] = true
	for {

		for c, i := range mapNode {

			spl := strings.Split(i, ",")
			for _, s := range spl {
				spl2 := strings.Split(strings.Trim(s, " ."), " ")
				newBag := (spl2[1] + " " + spl2[2])
				if finder[newBag] {
					finder[c] = true
				}

			}
		}
		if last == len(finder) {
			break
		}
		last = len(finder)
	}

	// substract shiny gold
	fmt.Println(len(finder) - 1)
}
