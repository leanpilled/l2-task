package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type anagramSet struct {
	set       map[string]struct{}
	firstWord string
}

func newAnagramSet(word string) *anagramSet {
	s := make(map[string]struct{})
	s[word] = struct{}{}

	return &anagramSet{
		set:       s,
		firstWord: word,
	}
}

func findAnagrams(dict []string) map[string][]string {
	res := map[string]anagramSet{}

	for _, word := range dict {
		word = strings.ToLower(word)
		sortedWord := sortString(word)
		s, ok := res[sortedWord]
		if ok {
			s.set[word] = struct{}{}
		} else {
			s = *newAnagramSet(word)
			res[sortedWord] = s
		}
	}

	anagrams := map[string][]string{}
	for key := range res {
		anagramsSlice := []string{}
		for w := range res[key].set {
			anagramsSlice = append(anagramsSlice, w)
		}
		if len(anagramsSlice) > 1 {
			sort.Strings(anagramsSlice)
			anagrams[res[key].firstWord] = anagramsSlice
		}
	}

	return anagrams
}

func sortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
