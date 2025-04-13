package main

import (
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
	crack := flag.Bool("x", false, "Crack input file")
	help := flag.Bool("h", false, "Print usage information")

	flag.Parse()

	if *help || (flag.NFlag() == 0) {
		printUsage()
		return
	}

	if !*encrypt && !*decrypt && !*frequency && !*crack {
		fmt.Println("Error: You must specify either -e (encrypt) or -d (decrypt) or -f frequency or -x crack.")
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
	} else if *crack {
		mode = "crack"
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
			os.Exit(1)
		}
		err = caesarcipher.StoreFrequencyFloat(output, freq)
		if err != nil {
			fmt.Printf("Error: cannot write output: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Frequency analysis complete")

	case "crack":
		engFreq, err := caesarcipher.LoadFrequencyFloat("./frequencies/english.json")
		if err != nil {
			fmt.Printf("Error: Failed to load english frequencies: %v\n", err)
			return
		}

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
