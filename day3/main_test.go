package main

import (
	"strings"
	"testing"
)

type testCase struct {
	mapTree string
	right   int
	down    int
	cnt     int
}

var testData = []testCase{

	{`.#
#.`, 1, 1, 0},
	{`#.
.#`, 1, 1, 2},
	{`#.
.#`, 3, 1, 2}, // repated slope
	{`...
...
...
...
...`, 3, 1, 0}, // repated slope, no trees
	{`..
..
.#`, 1, 2, 1},
	{`...
...
.#.
...
..#`, 1, 2, 2}, // repated slope, no trees
}

func TestSolve(t *testing.T) {

	for _, td := range testData {
		cnt := solve(strings.NewReader(td.mapTree), td.right, td.down)
		if cnt != td.cnt {
			t.Fatalf("%s should find %d tree(s), %d ", td.mapTree, td.cnt, cnt)
		}
	}
}
