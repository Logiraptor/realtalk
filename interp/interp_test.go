package interp

import "testing"
import "github.com/Logiraptor/realtalk/ast"
import "github.com/stretchr/testify/assert"

func TestInterpPrimitives(t *testing.T) {
	type testCase struct {
		input  ast.Expression
		output Expression
	}

	testCases := []testCase{
		{input: ast.Int(0), output: NewInt(0)},
	}

	for _, tc := range testCases {
		result, err := InterpExpression(tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.output, result)
	}
}

func TestNewInt(t *testing.T) {

}
