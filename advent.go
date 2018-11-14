package advent

import (
	"bufio"
	"log"
	"os"
)

// InputLines reads a file and returns the input split by lines. Exits with
// status 1 if the file doesn't exist.
func InputLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// MustGetArg returns the command line arg at position argIndex if it exists
// otherwise exits with status 1. Use this to get args from the command line
// when you want to avoid mysterious errors!
func MustGetArg(argIndex int) string {
	if len(os.Args) < argIndex+1 {
		log.Fatalf("No arg at position os.Args[%d]", argIndex)
	}
	return os.Args[argIndex]
}
