package s1768

import "testing"

type mergeAlternatelyTest struct {
	arg1, arg2, expected string
}

var mergeAlternatelyTests = []mergeAlternatelyTest{
	{"abc", "pqr", "apbqcr"},
	{"ab", "pqrs", "apbqrs"},
	{"abcd", "pq", "apbqcd"},
}

func TestMergeAlternately(t *testing.T) {
	for _, test := range mergeAlternatelyTests {
		output := MergeAlternately(test.arg1, test.arg2)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
