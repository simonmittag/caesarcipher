package caesarcipher

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"strings"
	"time"
	"unicode"
)

type CaesarCipher struct {
	offset    int
	Reference FrequencyFloat
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

func (c *CaesarCipher) Frequency(reader io.Reader) (Frequency, error) {
	scanner := bufio.NewScanner(reader)
	res := NewFrequency("Sampled-" + time.Now().Format(time.RFC3339))
	for scanner.Scan() {
		//we lowercase for analysis to reduce surface
		line := strings.ToLower(scanner.Text())
		res.Merge(c.FrequencyAnalysis(line))
	}
	err := scanner.Err()
	return *res, err
}

func (c *CaesarCipher) FrequencyAnalysis(line string) Frequency {
	f := NewFrequency("line")
	for _, char := range []rune(line) {
		f.Values[char]++
	}
	return *f
}

func (c *CaesarCipher) Crack(input io.Reader, output io.Writer) error {
	var inputBuffer bytes.Buffer
	if _, err := io.Copy(&inputBuffer, input); err != nil {
		return fmt.Errorf("failed to cache input: %w", err)
	}

	sse := math.MaxFloat64
	bestShift := 0

	for i := 1; i <= 26; i++ { //64kb is the max we read when finding the best offset
		inputReader := io.LimitReader(bytes.NewReader(inputBuffer.Bytes()), 2<<16)
		tempWriter := new(bytes.Buffer)

		c.offset = i
		if err := c.Shift(inputReader, tempWriter, true); err != nil {
			return fmt.Errorf("failed to shift with offset %d: %w", i, err)
		}

		shiftedInputReader := bytes.NewReader(tempWriter.Bytes())
		shiftFreq, err := c.Frequency(shiftedInputReader)
		if err != nil {
			return fmt.Errorf("failed to calculate shifted frequency for offset %d: %w", i, err)
		}

		shiftSse := c.Reference.Values.sumSquaredError(shiftFreq.ToFractions())
		if shiftSse < sse {
			// Update the best result.
			sse = shiftSse
			bestShift = i
		}
	}

	fmt.Printf("squared sum error: %f, shift: %d\n", sse, bestShift)
	c.offset = bestShift
	err := c.Shift(bytes.NewReader(inputBuffer.Bytes()), output, true)
	if err != nil {
		return fmt.Errorf("failed to write best result to output: %w", err)
	}
	return nil
}
