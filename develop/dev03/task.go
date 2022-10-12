package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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

type flags struct {
	filename  string
	column    int
	hasNumber bool
	isReverse bool
	isUnique  bool
}

func parseFlags() *flags {
	f := flags{}

	flag.StringVar(&f.filename, "f", "", "filename to sort")
	flag.IntVar(&f.column, "col", -1, "column to sort")
	flag.BoolVar(&f.hasNumber, "n", false, "sort as a number value")
	flag.BoolVar(&f.isReverse, "r", false, "reverse sort")
	flag.BoolVar(&f.isUnique, "u", false, "sort only unique strings")
	flag.Parse()

	return &f
}

func processSort(lines []string, f *flags) []string {
	if f.column >= 0 {
		sortByColumn(lines, f.column)
	} else {
		sort.Strings(lines)
	}

	if f.isReverse {
		reverse(lines)
	}

	if f.isUnique {
		lines = unique(lines)
	}

	return lines
}

func sortByColumn(lines []string, col int) []string {
	m := make(map[string]string)

	for i := 0; i < len(lines); i++ {
		splittedLine := strings.Split(lines[i], " ")
		m[splittedLine[col]] = lines[i]
	}

	keys := getKeys(m)
	sort.Strings(keys)

	sortedLines := make([]string, len(lines))
	for i, key := range keys {
		sortedLines[i] = m[key]
	}

	return sortedLines
}

func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func reverse(data []string) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func readFile(path string) []string {
	var rows []string

	f, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing file", err.Error())
		}
	}(f)

	if err != nil {
		log.Fatal("Error opening file", err.Error())
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}

	return rows
}

func writeFile(path string, data []string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal("Error creating file", err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error closing file", err.Error())
		}
	}(f)

	for i := 0; i < len(data); i++ {
		_, err := fmt.Fprintln(f, data[i])
		if err != nil {
			log.Fatal("Error writing to the file", err.Error())
		}
	}
}

func main() {
	const outputFile = "/Users/qnocks/GolandProjects/wildberries-internship/l2-project/develop/dev03/output.txt"

	flags := parseFlags()
	lines := readFile(flags.filename)
	sortedLines := processSort(lines, flags)
	writeFile(outputFile, sortedLines)
}
