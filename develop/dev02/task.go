package main

import (
	"fmt"
	"strconv"
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

func unbox(s string) string {
	if len(s) == 0 || unicode.IsDigit(rune(s[0])) {
		return ""
	}

	sb := strings.Builder{}
	for i, c := range s {
		if d, err := strconv.Atoi(string(c)); err == nil {
			for k := 1; k < d; k++ {
				sb.WriteString(string(s[i-1]))
			}
		} else {
			sb.WriteString(string(c))
		}
	}

	return sb.String()
}

func main() {
	fmt.Println(unbox(""))
}
