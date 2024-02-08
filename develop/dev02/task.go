package main

import (
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpackString(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	res := strings.Builder{}
	mult := 0
	runes := []rune(s)
	r := 1
	curr := runes[0]
	for ; r < len(runes); r++ {
		if unicode.IsDigit(runes[r]) {
			mult = mult*10 + int(runes[r]-'0')
		} else {
			if mult == 0 {
				mult = 1
			}
			res.WriteString(strings.Repeat(string(curr), mult))
			mult = 0
			if string(runes[r]) == "\\" {
				r += 1
			}
			curr = runes[r]
		}
	}
	if mult == 0 {
		mult = 1
	}
	res.WriteString(strings.Repeat(string(curr), mult))
	return res.String()
}
