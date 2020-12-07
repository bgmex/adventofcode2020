package main

import (
	"testing"
)

func TestCheckYR(t *testing.T) {
	type TestCase struct {
		s     string
		min   int
		max   int
		valid bool
	}

	var TestCases []TestCase = []TestCase{
		{"1900",
			1920,
			2002,
			false},
		{"2002",
			1920,
			2002,
			true},
		{"2003",
			1920,
			2002,
			false},
		{"1920",
			1920,
			2002,
			true},
		{"1920",
			2010,
			2020,
			false},
	}

	for _, tc := range TestCases {
		if checkYR(tc.s, tc.min, tc.max) != tc.valid {
			t.Fatalf("%s is valid", tc.s)
		}
	}

}

func TestCheckHGT(t *testing.T) {
	type TestCase struct {
		s     string
		valid bool
	}

	var TestCases []TestCase = []TestCase{
		{"60in",
			true},
		{"60cm",
			false},
		{"190cm",
			true},
		{"190in",
			false},
		{"150cm",
			true},
		{"193cm",
			true},
		{"59in",
			true},
		{"76in",
			true},
		{"76cmin",
			false},
		{"76incm",
			false},
		{"in76",
			false},
		{"76",
			false},
	}
	for _, tc := range TestCases {
		if checkHGT(tc.s) != tc.valid {
			t.Fatalf("%s is valid", tc.s)
		}
	}
}

func TestCheckHCL(t *testing.T) {
	type TestCase struct {
		s     string
		valid bool
	}

	var TestCases []TestCase = []TestCase{
		{"#123abc",
			true},
		{"123abz",
			false},
		{"123abc",
			false},
		{"0x12345",
			false},
		{"123abcd",
			false},
		{"#0x1234",
			false},
	}
	for _, tc := range TestCases {
		if checkHCL(tc.s) != tc.valid {
			t.Fatalf("%s is valid", tc.s)
		}
	}

}
func TestCheckECL(t *testing.T) {
	type TestCase struct {
		s     string
		valid bool
	}

	var TestCases []TestCase = []TestCase{
		{"brn",
			true},
		{"wat",
			false},
		{"brnn",
			false},
	}
	for _, tc := range TestCases {
		if checkECL(tc.s) != tc.valid {
			t.Fatalf("%s is valid", tc.s)
		}
	}

}
func TestCheckPID(t *testing.T) {
	type TestCase struct {
		s     string
		valid bool
	}

	var TestCases []TestCase = []TestCase{
		{"000000001",
			true},
		{"0123456789",
			false},
	}
	for _, tc := range TestCases {
		if checkPID(tc.s) != tc.valid {
			t.Fatalf("%s is valid", tc.s)
		}
	}

}
