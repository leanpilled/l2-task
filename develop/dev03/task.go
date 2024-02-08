package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type manSort struct {
	Lines   []string
	Column  int
	Sep     string
	Reverse bool
	Unique  bool
}

func (b manSort) Len() int      { return len(b.Lines) }
func (b manSort) Swap(i, j int) { b.Lines[i], b.Lines[j] = b.Lines[j], b.Lines[i] }

func (b manSort) Less(i, j int) bool {
	fieldsI := strings.Split(b.Lines[i], " ")
	fieldsJ := strings.Split(b.Lines[j], " ")

	if b.Column < len(fieldsI) && b.Column < len(fieldsJ) {
		if b.Reverse {
			return fieldsI[b.Column] > fieldsJ[b.Column]
		}
		return fieldsI[b.Column] < fieldsJ[b.Column]
	}
	return false
}

func main() {
	filePath := flag.String("file", "", "path to the file")
	columnToSortBy := flag.Int("k", 0, "column number (0-based) to sort by")
	reverse := flag.Bool("r", false, "reverse sorting")
	unique := flag.Bool("u", false, "unique lines")

	flag.Parse()

	lines, err := readLines(*filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	byColumn := manSort{
		Lines:   lines,
		Column:  *columnToSortBy,
		Reverse: *reverse,
		Unique:  *unique,
	}

	sort.Sort(byColumn)

	if *unique {
		unique := map[string]struct{}{}
		for _, line := range byColumn.Lines {
			if _, ok := unique[line]; !ok {
				unique[line] = struct{}{}
				fmt.Println(line)
			}
		}
	} else {
		for _, line := range byColumn.Lines {
			fmt.Println(line)
		}
	}
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
