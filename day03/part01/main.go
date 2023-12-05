package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func getLines() []string {
	dir, _ := os.Getwd()
	pathToInputFile := path.Join(dir, "day03", "input.txt")
	data, err := os.ReadFile(pathToInputFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines
}

func isSymbol(row int, col int, lines []string) bool {
	n := len(lines)
	m := len(lines[0])
	if !(0 <= row && row < n && 0 <= col && col < m) {
		return false
	}
	reSymbol := regexp.MustCompile(`[^0-9.]`)
	isSymbol := reSymbol.MatchString(string(lines[row][col]))
	return isSymbol
}

func main() {
	lines := getLines()
	m := len(lines[0])
	answer := 0

	for i, line := range lines {
		start := 0
		j := 0

		for j < m {
			start = j
			numStr := ""
			for j < m && ('0' <= line[j] && line[j] <= '9') {
				numStr += string(line[j])
				j += 1
			}
			if numStr == "" {
				j += 1
				continue
			}
			num, _ := strconv.Atoi(numStr)
			if isSymbol(i, start-1, lines) || isSymbol(i, j, lines) {
				answer += num
				continue
			}
			for k := start - 1; k <= j; k++ {
				if isSymbol(i-1, k, lines) || isSymbol(i+1, k, lines) {
					answer += num
					break
				}
			}
		}
	}
	fmt.Print(answer)
}
