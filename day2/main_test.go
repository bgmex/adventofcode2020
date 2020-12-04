package main

import (
	"testing"
)

type testCase struct {
	line  string
	valid bool
}

var testData = []testCase{

	{"1-3 a: sdf", false},     // no chars
	{"1-3 a: asdf", true},     // minimum chars
	{"1-3 a: aasdf", true},    // in between min and max chars
	{"1-3 a: aaasdf", true},   // maximum chars
	{"1-3 a: aaaasdf", false}, // too many chars
}

func TestCheck(t *testing.T) {

	for _, td := range testData {
		if check(td.line) != td.valid {
			t.Fatalf("%s should not be %t", td.line, td.valid)
		}
	}
}
