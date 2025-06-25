package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Prints a promt to stdout and returns int, taken from stdin
func GetInt(promt string) int {
	fmt.Print(promt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading error. Enter one integer and press 'Enter'\n")
		return GetInt(promt)
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input error. Enter one integer and press 'Enter'\n")
		return GetInt(promt)
	}

	return number
}

// Prints a promt to stdout and returns first string, taken from stdin
func GetString(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading error. Enter one string and press 'Enter'\n")
		return GetString(promt)
	}
	input = strings.TrimSpace(input)
	return input
}
