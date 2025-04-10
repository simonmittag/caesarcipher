package caesarcipher

import (
	"testing"
)

func TestShiftWithOffset(t *testing.T) {
	cc := NewCaesarCipher(0)

	tests := []struct {
		input    string
		shift    int
		expected string
	}{
		{"ABC", 3, "DEF"},
		{"XYZ", 3, "ABC"},
		{"abc", 3, "def"},
		{"xyz", 3, "abc"},
		{"ABC", -3, "XYZ"},
		{"XYZ", -3, "UVW"},
		{"abc", -3, "xyz"},
		{"xyz", -3, "uvw"},
		{"Hello, World!", 5, "Mjqqt, Btwqi!"},
		{"", 10, ""},
		{"123", 5, "123"},
	}

	for _, test := range tests {
		result := cc.ShiftWithOffset(test.input, test.shift)
		if result != test.expected {
			t.Errorf("ShiftWithOffset(%q, %d) = %q; want %q", test.input, test.shift, result, test.expected)
		}
	}
}
