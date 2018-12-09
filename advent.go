package advent

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// BigInt is the max int.
var BigInt = 1<<63 - 1

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

// StringsToInts takes an array of lines from a file read and safely converts them
// to an array of ints.
func StringsToInts(lines []string) []int {
	ints := make([]int, 0, len(lines))
	for lineno, line := range lines {
		i, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatalf("Unable to parse int at position %v\n%v", lineno+1, err)
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

// StripNonDigits replaces all non-digit characters in line with spaces.
func StripNonDigits(line string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(line, " ")
}

// Split slices line into substrings separated by any amount of whitespace.
func Split(line string) []string {
	line = strings.TrimSpace(line)
	re := regexp.MustCompile("[[:space:]]+")
	return re.Split(line, -1)
}

// MaxInts return the max in in xs
func MaxInts(xs ...int) int {
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

// Max returns the greater of a and b.
func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// Min returns the lesser of a and b.
func Min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Abs returns |x|.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// SumInts return the sum of the ints in xs.
func SumInts(xs ...int) int {
	sum := 0
	for i := range xs {
		sum += xs[i]
	}
	return sum
}
