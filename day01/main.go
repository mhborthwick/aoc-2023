package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func getInputs(inputFile string) []string {
	dir, _ := os.Getwd()
	pathToInputFile := path.Join(dir, "day01", inputFile)
	file, _ := os.Open(pathToInputFile)
	b, _ := io.ReadAll(file)
	defer func() {
		_ = file.Close()
	}()
	return strings.Split(string(b), "\n")
}

func getFirstNumber(re *regexp.Regexp, input string) string {
	for _, i := range input {
		if re.MatchString(string(i)) {
			return string(i)
		}
	}
	return ""
}

func getLastNumber(re *regexp.Regexp, input string) string {
	for i := len(input) - 1; i >= 0; i-- {
		if re.MatchString(string(input[i])) {
			return string(input[i])
		}
	}
	return ""
}

func main() {
	total := 0
	re := regexp.MustCompile("[0-9]")
	inputs := getInputs("./input.txt")
	for _, i := range inputs {
		first := getFirstNumber(re, i)
		last := getLastNumber(re, i)
		merge := first + last
		s, _ := strconv.Atoi(merge)
		total += s
	}
	fmt.Print(total)
}
