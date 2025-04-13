package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/simonmittag/caesarcipher"
	"os"
)

const Version = "0.1.1"

func main() {

	var mode string
	inputFile := flag.String("i", "", "Path to the input text file (required)")
	outputFile := flag.String("o", "", "Path to the output text file (required)")
	offset := flag.Int("s", 0, "Offset for the Caesar cipher (default: 0)")
	encrypt := flag.Bool("e", false, "Encrypt the input file")
	decrypt := flag.Bool("d", false, "Decrypt the input file")
	frequency := flag.Bool("f", false, "Frequency analysis input file")
	help := flag.Bool("h", false, "Print usage information")

	flag.Parse()

	if *help || (flag.NFlag() == 0) {
		printUsage()
		return
	}

	if !*encrypt && !*decrypt && !*frequency {
		fmt.Println("Error: You must specify either -e (encrypt) or -d (decrypt) or -f frequency.")
		printUsage()
		os.Exit(1)
	}
	if *inputFile == "" {
		fmt.Println("Error: Input file (-i flag) is required.")
		printUsage()
		os.Exit(1)
	}
	if *outputFile == "" {
		fmt.Println("Error: Output file (-o flag) is required.")
		printUsage()
		os.Exit(1)
	}

	if *encrypt {
		mode = "encrypt"
	} else if *decrypt {
		mode = "decrypt"
	} else if *frequency {
		mode = "frequency"
	}

	input, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error: Could not open input file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	output, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error: Could not create output file: %v\n", err)
		os.Exit(1)
	}
	defer output.Close()

	cipher := caesarcipher.NewCaesarCipher(*offset)

	switch mode {
	case "encrypt":
		err = cipher.Shift(input, output, false) // false indicates encryption
	case "decrypt":
		err = cipher.Shift(input, output, true) // true indicates decryption
	case "frequency":
		freq, err := cipher.Frequency(input)
		if err != nil {
			fmt.Printf("Error: Frequency analysis failed: %v\n", err)
			return
		}

		jsonData, err := json.MarshalIndent(freq.ToFractions(), "", "  ")
		if err != nil {
			fmt.Printf("Error: Failed to encode frequency analysis to JSON: %v\n", err)
			return
		}

		_, err = output.Write(jsonData)
		if err != nil {
			fmt.Printf("Error: Failed to write JSON output to file: %v\n", err)
			return
		}

		fmt.Println("Frequency analysis complete")
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Output written to %s\n", *outputFile)
}

// printUsage prints the usage instructions
func printUsage() {
	fmt.Println("🌿 Caesar " + Version)
	fmt.Println("Usage: caesar [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
