package caesarcipher

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type CaesarCipher struct {
	offset int
}

// NewCaesarCipher creates a new CaesarCipher with the given offset.
func NewCaesarCipher(offset int) *CaesarCipher {
	return &CaesarCipher{offset: offset}
}

// Shift handles both encryption and decryption based on a flag.
// Pass `decrypt = true` for decryption (reverses the offset).
func (c *CaesarCipher) Shift(reader io.Reader, writer io.Writer, decrypt bool) error {
	// Adjust the offset for decryption
	shift := c.offset
	if decrypt {
		shift = -c.offset
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		resultLine := c.ShiftWithOffset(line, shift)
		_, err := fmt.Fprintln(writer, resultLine)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

// ShiftWithOffset performs the character rotation (shift) based on the given offset.
func (c *CaesarCipher) ShiftWithOffset(input string, shift int) string {
	shift = shift % 26 // Ensure shift is within bounds (0-25).
	runes := []rune(input)

	for i, r := range runes {
		// we skip everything that's not in the roman alphabet
		if unicode.IsLetter(r) {
			base := 'A'
			if unicode.IsLower(r) {
				base = 'a'
			}

			relativePosition := r - base

			//a few things happen here. We shift to a new position.
			//but in case it's negative, we add 26, then mod 26
			//so we overflow back into the bounds of the same alphabet
			newPosition := relativePosition + rune(shift)
			overflowPosition := (newPosition + 26) % 26
			runes[i] = base + overflowPosition
		}
	}

	return string(runes)
}
