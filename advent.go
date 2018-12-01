package advent

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

// LinesToInts takes an array of lines from a file read and safely converts them
// to an array of ints.
func LinesToInts(lines []string) []int {
	ints := make([]int, 0, len(lines))
	for lineno, line := range lines {
		i, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatalf("Unable to parse int on line number %v\n%v", lineno+1, err)
		}
		ints = append(ints, i)
	}
	return ints
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
