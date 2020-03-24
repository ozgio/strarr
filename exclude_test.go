package strarr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExclude(t *testing.T) {
	tests := []struct {
		a        []string
		b        []string
		expected []string
	}{
		{[]string{"a", "b"}, []string{"c", "d"}, []string{"a", "b"}},
		{[]string{"a", "b"}, []string{"a", "d"}, []string{"b"}},
		{[]string{"a", "b"}, []string{"a", "a"}, []string{"b"}},
		{[]string{"a", "b"}, []string{"a", "b"}, []string{}},
		{[]string{"a", "b"}, []string{}, []string{"a", "b"}},
		{[]string{}, []string{"a", "b"}, []string{}},
		{[]string{}, []string{}, []string{}},
	}

	for i, test := range tests {
		res := Exclude(test.a, test.b)
		assert.Equal(t, test.expected, res, "%d: not the expected value", i)
	}
}
