# Caesar Cipher

Monoalphabetic substitution is one of the simplest and oldest encryption techniques, shifting the alphabet by a fixed number of positions to encode or decode text. This project provides a lightweight command-line tool, written in Go, to encrypt and decrypt text files using the Caesar cipher.

## Getting Started

### Prerequisites

This project requires **Go** to be installed. To set up Go, you can use the following command on macOS:

```bash
brew install go
```

### Installation

To install the Caesar cipher command-line tool, use:

```bash
go install github.com/simonmittag/caesarcipher/cmd/caesar
```

This will make the `caesar` command available globally.

## Usage

The `caesar` command-line tool supports encrypting and decrypting text files using the Caesar cipher. Below are the available options and examples to get you started:

### Options

```bash
🌿 CaesarCipher 0.1.0
Usage: caesar [options]

Options:
  -e                Encrypt the input file.
  -d                Decrypt the input file.
  -s int            Offset for the Caesar cipher (default: 0).
  -i string         Path to the input text file (required).
  -o string         Path to the output text file (required).
  -h                Print usage information.
```

### Examples

#### Encrypt a File

To encrypt the content of `mary.txt` using a shift of `3` and save the result to `mary_encrypted.txt`:

```bash
caesar -e -s 3 -i mary.txt -o mary_encrypted.txt
```