package strarr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		a        []string
		expected []string
	}{
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "b"}, []string{"a", "b"}},
		{[]string{"b", "b"}, []string{"b"}},
		{[]string{"a"}, []string{"a"}},
		{[]string{}, []string{}},
	}

	for i, test := range tests {
		res := Unique(test.a)
		assert.Equal(t, test.expected, res, "%d: not the expected value", i)
	}
}
