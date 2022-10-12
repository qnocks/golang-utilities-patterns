package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	A, B, C           int
	c, i, v, F, n     bool
	pattern, filename string
}

func grep() {
	f, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := readFile(f.filename)
	if err != nil {
		log.Fatalln(err)
	}
	if f.C > 0 {
		f.A = f.C
		f.B = f.C
	}

	res, key := findAllStrings(f, file)

	if f.c {
		fmt.Println(len(res))
		return
	}

	printResult(f, res, key)
}

func printResult(f *flags, res map[int]string, key []int) {
	for _, v := range key {
		if f.n {
			fmt.Print(v, ":", res[v], "\n")
		} else {
			fmt.Println(res[v])
		}
	}
}

func findAllStrings(f *flags, file []string) (map[int]string, []int) {
	result := make(map[int]string)
	key := make([]int, 0)
	write := false
	for i, v := range file {
		w := v
		if f.i {
			w = strings.ToLower(w)
			f.pattern = strings.ToLower(f.pattern)
		}
		if f.F {
			if strings.Contains(w, f.pattern) {
				result[i+1] = v
				key = append(key, i+1)
				write = true
			}
		} else {
			check, err := regexp.MatchString(f.pattern, w)
			if err != nil {
				log.Fatal(err)
			}
			if f.v {
				if !check {
					result[i+1] = v
					key = append(key, i+1)
					write = true
				}
			} else {
				if check {
					result[i+1] = v
					key = append(key, i+1)
					write = true
				}
			}
			if (f.B > 0 || f.A > 0) && write && !f.c {
				j := i + 1
				for n := f.A; n > 0; n-- {
					if j < len(file) {
						result[j+1] = file[j]
						key = append(key, j+1)
						j++
					}
				}
				j = i - 1
				for n := f.B; n > 0; n-- {
					result[j+1] = file[j]
					key = append(key, j+1)
					j--
				}
				write = false
			}
		}
	}
	sort.Ints(key)
	return result, key
}

func readFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buffer, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	result := strings.Split(string(buffer), "\n")
	return result, nil
}

func parseFlags() (*flags, error) {
	f := flags{}

	flag.IntVar(&f.A, "A", -1, "\"after\" print +N strings after match")
	flag.IntVar(&f.B, "B", -1, "\"before\" print +N strings before match")
	flag.IntVar(&f.C, "C", -1, "\"context\" (A+B) print ±N strings around match")
	flag.BoolVar(&f.c, "c", false, "\"count\" (strings count)")
	flag.BoolVar(&f.i, "i", false, "\"ignore-case\" (ignore upper/lower cases)")
	flag.BoolVar(&f.v, "v", false, "\"invert\" (instead of match, exclude)")
	flag.BoolVar(&f.F, "F", false, "\"fixed\", exact string match, not a pattern")
	flag.BoolVar(&f.n, "n", false, "\"line num\", print string number")

	flag.Parse()
	if len(flag.Args()) != 2 {
		return nil, errors.New("enter string for found and filename")
	}
	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	f.pattern = pattern
	f.filename = filename
	return &f, nil
}

func main() {
	grep()
}
