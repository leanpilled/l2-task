package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fieldsFlag := flag.String("f", "", "выбрать поля (колонки)")
	delimiterFlag := flag.String("d", "\t", "использовать другой разделитель")
	separatedFlag := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	delimiter := *delimiterFlag

	fields := make(map[int]bool)
	if *fieldsFlag != "" {
		for _, fieldStr := range strings.Split(*fieldsFlag, ",") {
			field, err := strconv.Atoi(fieldStr)
			if err != nil {
				panic(err)
			}
			fields[field] = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separatedFlag && !strings.Contains(line, delimiter) {
			continue
		}

		columns := strings.Split(line, delimiter)

		var selectedColumns []string
		for i, col := range columns {
			if len(fields) == 0 || fields[i+1] {
				selectedColumns = append(selectedColumns, col)
			}
		}

		fmt.Println(strings.Join(selectedColumns, delimiter))
	}
}
