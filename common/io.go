package common

import (
	"bufio"
	"os"
)

// ReadFileAsInts reads all lines from the given path and returns them in a slice of ints.
// If an error occurs, the function will panic.
func ReadFileAsInts(path string) []int {
	return readIntsWithScanner(path, bufio.ScanLines)
}

// ReadFileAsStrings reads all lines from the given path and returns them in a slice of strings.
// If an error occurs, the function will panic.
func ReadFileAsStrings(path string) []string {
	return readStringsWithScanner(path, bufio.ScanLines)
}

// ReadCsvAsInts reads all data from the given path and returns an int slice
// containing comma-delimited parts.  If an error occurs, the function will panic.
func ReadCsvAsInts(path string) []int {
	return readIntsWithScanner(path, scanByCommas)
}

// readIntsWithScanner uses a bufio.Scanner to read ints from the file at
// the given path, splitting using the given bufio.SplitFunc. If an error
// occurs at any point, the function will panic.
func readIntsWithScanner(path string, splitFunc bufio.SplitFunc) []int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var parts []int
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		parts = append(parts, MustAtoi(scanner.Text()))
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return parts
}

// readStringsWithScanner uses a bufio.Scanner to read strings from the file at
// the given path, splitting using the given bufio.SplitFunc. If an error
// occurs at any point, the function will panic.
func readStringsWithScanner(path string, splitFunc bufio.SplitFunc) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var parts []string
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		parts = append(parts, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return parts
}

// scanByCommas is a split function for a Scanner that returns each
// comma-separated section of text, with surrounding spaces deleted.
// The definition of space is set by unicode.IsSpace.
func scanByCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		if data[start] != ' ' {
			break
		}
	}

	// Scan until comma, marking end of word.
	for i := start; i < len(data); i++ {
		if data[i] == ',' || data[i] == ' ' || data[i] == '\r' || data[i] == '\n' {
			return i + 1, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
