package test

import "testing"
import "github.com/stretchr/testify/assert"

type endToEndResult struct {
	stdout string
}

func runEndToEnd(source string) endToEndResult {
	return endToEndResult{
		stdout: "1",
	}
}

func TestEndToEndLiterals(t *testing.T) {
	type testCase struct {
		input  string
		output string
	}

	testCases := []testCase{
		{input: "System.print(1)", output: "1"},
		{input: "System.print(\"Hello, World!\")", output: "Hello, World!"},
	}

	for _, tc := range testCases {
		result := runEndToEnd(tc.input)
		assert.Equal(t, result.stdout, tc.output)
	}
}
