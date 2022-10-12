package main

import (
	"fmt"
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

func search(words *[]string) *map[string][]string {
	result := make(map[string][]string)
	toLower(*words)
	sort.Strings(*words)

	for _, w := range *words {
		isStored := false
		for k, v := range result {
			if isAnagram(k, w) {
				result[k] = append(v, w)
				isStored = true
				break
			}
		}
		if !isStored {
			result[w] = append(result[w], w)
		}
	}

	for k, v := range result {
		if len(v) <= 1 {
			delete(result, k)
		}
	}

	return &result
}

func isAnagram(first, second string) bool {
	return sortChars(first) == sortChars(second)
}

func sortChars(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func toLower(data []string) {
	for i, x := range data {
		data[i] = strings.ToLower(x)
	}
}

func main() {
	words := []string{"Тяпка", "пятак", "лИсток", "столик", "пятка", "слиток"}
	m := search(&words)
	fmt.Println(m)
}
