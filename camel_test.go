package phrase

import (
	"fmt"
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
		tt := tt
		t.Run(fmt.Sprintf("S=%s,Sep=%s", tt.s, tt.sep), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, ToCamel(tt.s, tt.sep))
		})
	}
}
