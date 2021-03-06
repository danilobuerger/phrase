// Copyright (c) 2016 Danilo Bürger <info@danilobuerger.de>

package phrase

import (
	"fmt"
	"testing"
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
			got := ToCamel(tt.s, tt.sep)
			if got != tt.expected {
				t.Errorf("ToCamel(%s, %s) => %s, expected %s", tt.s, tt.sep, got, tt.expected)
			}
		})
	}
}
