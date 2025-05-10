package format_test

import (
	"testing"

	"github.com/amadejkastelic/spar-api/internal/format"
)

func TestFormatFloat(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		sep      string
		expected string
	}{
		{"Zero", 0, ",", "0,00"},
		{"Comma", 1234.56, ",", "1234,56"},
		{"Dot", 1234.5678, ".", "1234.57"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := format.FormatFloat(tc.input, tc.sep)
			if result != tc.expected {
				t.Errorf("FormatFloat(%f, %q) = %q; want %q", tc.input, tc.sep, result, tc.expected)
			}
		})
	}
}
