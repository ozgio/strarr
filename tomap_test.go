package strarr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var empty = struct{}{}

func TestToMap(t *testing.T) {
	tests := []struct {
		a        []string
		expected map[string]struct{}
	}{
		{[]string{"a", "b"}, map[string]struct{}{"a": empty, "b": empty}},
		{[]string{"a", "", ""}, map[string]struct{}{"a": empty, "": empty}},
		{[]string{"a", "a"}, map[string]struct{}{"a": empty}},
		{[]string{}, map[string]struct{}{}},
	}

	for i, test := range tests {
		res := ToMap(test.a)
		assert.Equal(t, test.expected, res, "%d: not the expected value", i)
	}
}
