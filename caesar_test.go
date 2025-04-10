package caesarcipher

import (
	"bytes"
	"strings"
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

func TestShift(t *testing.T) {
	tests := []struct {
		name      string
		offset    int
		input     string
		decrypt   bool
		want      string
		expectErr bool
	}{
		{
			name:   "Encrypt multiline text",
			offset: 2,
			input: `Hello, World!
This is a test.
Caesar cipher!`,
			decrypt: false,
			want: `Jgnnq, Yqtnf!
Vjku ku c vguv.
Ecguct ekrjgt!`,
		},
		{
			name:   "Decrypt multiline text",
			offset: 2,
			input: `Jgnnq, Yqtnf!
Vjku ku c vguv.
Ecguct ekrjgt!`,
			decrypt: true,
			want: `Hello, World!
This is a test.
Caesar cipher!`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			writer := &bytes.Buffer{}

			cipher := NewCaesarCipher(tc.offset)
			err := cipher.Shift(reader, writer, tc.decrypt)

			if tc.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			got := writer.String()
			if strings.TrimSpace(got) != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
