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
	pathToInputFile := path.Join(dir, "day02", inputFile)
	file, _ := os.Open(pathToInputFile)
	b, _ := io.ReadAll(file)
	defer func() {
		_ = file.Close()
	}()
	return strings.Split(string(b), "\n")
}

func getValues(re *regexp.Regexp, input string) []string {
	matches := re.FindAllString(input, -1)
	return matches
}

func isPossible(input []string) bool {
	isValid := true
	reNumber := regexp.MustCompile(`\d+`)
	reText := regexp.MustCompile(`(?:blue|red|green)`)
	for _, v := range input {
		n := reNumber.FindString(v)
		t := reText.FindString(v)
		s, _ := strconv.Atoi(n)
		if t == "red" && s > 12 {
			isValid = false
			break
		}
		if t == "green" && s > 13 {
			isValid = false
			break
		}
		if t == "blue" && s > 14 {
			isValid = false
			break
		}
	}
	return isValid
}

func main() {
	total := 0
	inputs := getInputs("input.txt")
	re := regexp.MustCompile(`\d+\s(?:blue|red|green)`)
	for i, v := range inputs {
		s := getValues(re, v)
		possible := isPossible(s)
		if possible {
			total += i + 1
		}
	}
	fmt.Print(total)
}
