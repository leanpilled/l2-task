package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения OK
-B - "before" печатать +N строк до совпадения OK
-C - "context" (A+B) печатать ±N строк вокруг совпадения OK
-c - "count" (количество строк) OK
-i - "ignore-case" (игнорировать регистр) OK
-v - "invert" (вместо совпадения, исключать) OK
-F - "fixed", точное совпадение со строкой, не паттерн OK
-n - "line num", печатать номер строки OK

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type manGrep struct {
	filename   string
	pattern    string
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func (f manGrep) grep() error {
	if f.ignoreCase {
		f.pattern = strings.ToLower(f.pattern)
	}
	re := regexp.MustCompile(f.pattern)
	lines, err := readLines(f.filename)
	if err != nil {
		return err
	}

	matches := []int{}

	for i, line := range lines {
		if ok := f.match(line, re); ok {
			matches = append(matches, i)
		}
	}

	if f.after != 0 || f.before != 0 {
		matches = uniteIntervals(matches, f.before, f.after, len(lines))
	} else if f.context != 0 {
		matches = uniteIntervals(matches, f.context, f.context, len(lines))
	}

	if f.count {
		fmt.Println(len(matches))
	} else if f.invert {
		m := map[int]struct{}{}
		for _, i := range matches {
			m[i] = struct{}{}
		}
		for i, line := range lines {
			if _, ok := m[i]; !ok {
				fmt.Println(line)
			}
		}
	} else if f.lineNum {
		fmt.Println(matches)
	} else {
		for _, i := range matches {
			fmt.Println(lines[i])
		}
	}

	return nil
}

func uniteIntervals(matches []int, before int, after int, max int) []int {
	intervals := [][]int{}
	for _, m := range matches {
		r := m + after
		if m+after >= max {
			r = max - 1
		}
		l := m - before
		if m-before < 0 {
			l = 0
		}
		intervals = append(intervals, []int{l, r})
	}

	res := []int{}
	for _, interval := range intervals {
		if len(res) == 0 || res[len(res)-1] < interval[0] {
			for i := interval[0]; i <= interval[1]; i++ {
				res = append(res, i)
			}
		} else {
			for i := res[len(res)-1] + 1; i <= interval[1]; i++ {
				res = append(res, i)
			}
		}
	}

	return res
}

func (f manGrep) match(line string, re *regexp.Regexp) bool {
	if f.fixed {
		return f.pattern == line
	}
	if f.ignoreCase {
		return re.MatchString(strings.ToLower(line))
	}
	return re.MatchString(line)
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	var grep manGrep

	flag.StringVar(&grep.filename, "file", "", "File name")
	flag.StringVar(&grep.pattern, "reg", ".*", "Regexp")
	flag.IntVar(&grep.after, "A", 0, "Print N lines after each match")
	flag.IntVar(&grep.before, "B", 0, "Print N lines before each match")
	flag.IntVar(&grep.context, "C", 0, "Print N lines of output context")
	flag.BoolVar(&grep.count, "c", false, "Count of matching lines only")
	flag.BoolVar(&grep.ignoreCase, "i", false, "Ignore case distinctions")
	flag.BoolVar(&grep.invert, "v", false, "Invert the sense of matching")
	flag.BoolVar(&grep.fixed, "F", false, "Fixed, exact matching")
	flag.BoolVar(&grep.lineNum, "n", false, "Display line numbers")

	flag.Parse()

	err := grep.grep()
	if err != nil {
		fmt.Println(err)
	}
}
