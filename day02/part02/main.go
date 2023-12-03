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

func getPower(input []string) int {
	reNumber := regexp.MustCompile(`\d+`)
	reText := regexp.MustCompile(`(?:blue|red|green)`)
	m := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	for _, v := range input {
		n := reNumber.FindString(v)
		t := reText.FindString(v)
		s, _ := strconv.Atoi(n)
		if m[t] < s {
			m[t] = s
		}
	}
	power := m["red"] * m["blue"] * m["green"]
	return power
}

func main() {
	total := 0
	inputs := getInputs("input.txt")
	re := regexp.MustCompile(`\d+\s(?:blue|red|green)`)
	for _, v := range inputs {
		s := getValues(re, v)
		power := getPower(s)
		total += power
	}
	fmt.Print(total)
}
