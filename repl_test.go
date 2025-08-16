package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello there world  ",
			expected: []string{"hello", "there", "world"},
		},
		{
			input:    "  HeLlo tHere World 111 ",
			expected: []string{"hello", "there", "world", "111"},
		},
		{
			input:    "11 222 3   444",
			expected: []string{"11", "222", "3", "444"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cases with different length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words dont match")
			}
		}
	}
}
