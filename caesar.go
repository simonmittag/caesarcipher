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

func NewCaesarCipher(offset int) *CaesarCipher {
	return &CaesarCipher{offset: offset}
}

func (c *CaesarCipher) Shift(reader io.Reader, writer io.Writer, decrypt bool) error {
	//decryption shifts left, encryption shifts right
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

func (c *CaesarCipher) ShiftWithOffset(input string, shift int) string {
	shift = shift % 26 // Ensure shift is within bounds (0-25).
	runes := []rune(input)

	for i, letterAsNumber := range runes {
		// we skip everything that's not in the roman alphabet
		if unicode.IsLetter(letterAsNumber) {
			alphabetBase := 'A'
			if unicode.IsLower(letterAsNumber) {
				alphabetBase = 'a'
			}

			offBase := letterAsNumber - alphabetBase

			//a few things happen here. We shift to a new position.
			newPosition := offBase + rune(shift)

			//but in case it's negative, we add 26, then mod 26
			//so we overflow back into the bounds of the same alphabet
			overflowPosition := (newPosition + 26) % 26

			//and now we replace the value
			runes[i] = alphabetBase + overflowPosition
		}
	}

	//at the end all numbers are converted back to a string with proper letters.
	return string(runes)
}
