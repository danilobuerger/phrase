package phrase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamel(t *testing.T) {
	tests := []struct {
		s, sep   string
		expected string
	}{
		{"", "", ""},
		{"", "-", ""},
		{"a", "", "A"},
		{"a", "-", "A"},
		{"abc", "", "Abc"},
		{"abc", "-", "Abc"},
		{"abc-def", "", "Abc-Def"},
		{"abc-def", "-", "AbcDef"},
		{"abc def", " ", "AbcDef"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, ToCamel(tt.s, tt.sep))
	}
}
