package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	test := map[string]struct {
		input    string
		expected []string
	}{
		"simple":                      {input: "hello   world", expected: []string{"hello", "world"}},
		"trailing white space":        {input: "hello world ", expected: []string{"hello", "world"}},
		"white space at the begining": {input: "   hello world", expected: []string{"hello", "world"}},
		"empty string":                {input: "", expected: []string{}},
		"strings with number":         {input: "  123 45   8590  ", expected: []string{"123", "45", "8590"}},
		"special characters":          {input: " @! % ^ *", expected: []string{"@!", "%", "^", "*"}},
		"mixed cases":                 {input: "Hello wOrLd", expected: []string{"hello", "world"}},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Fatalf("%v", diff)
			}
		})
	}
}
