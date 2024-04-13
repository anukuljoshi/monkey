package evaluator

import (
	"testing"

	"github.com/anukuljoshi/monkey/lexer"
	"github.com/anukuljoshi/monkey/object"
	"github.com/anukuljoshi/monkey/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj is not Integer got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("result.Value: expected=%d, got=%d",
			expected, result.Value)
		return false
	}
	return true
}

// integer
func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj is not Boolean got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("result.Value: expected=%t, got=%t",
			expected, result.Value)
		return false
	}
	return true
}

// boolean
func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}
