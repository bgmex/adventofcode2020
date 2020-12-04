package main

import (
	"testing"
)

type testCase struct {
	line  string
	valid bool
}

var testData = []testCase{

	{"1-3 a: sdf", false},    // no match
	{"1-3 a: asdf", true},    // first pos
	{"1-3 a: rsaf", true},    // second pos
	{"1-3 a: aaasdf", false}, // both positions
}

func TestCheck(t *testing.T) {

	for _, td := range testData {
		if check(td.line) != td.valid {
			t.Fatalf("%s should not be %t", td.line, td.valid)
		}
	}
}
