package strarr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		a        []string
		b        []string
		expected []string
	}{
		{[]string{"a", "b"}, []string{"c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "b"}, []string{"a", "d"}, []string{"a", "b", "d"}},
		{[]string{"a", "b"}, []string{"a", "a"}, []string{"a", "b"}},
		{[]string{"a", "b"}, []string{}, []string{"a", "b"}},
		{[]string{}, []string{"a", "b"}, []string{"a", "b"}},
		{[]string{}, []string{}, []string{}},
	}

	for i, test := range tests {
		res := Merge(test.a, test.b)
		assert.Equal(t, test.expected, res, "%d: not the expected value", i)
	}
}
