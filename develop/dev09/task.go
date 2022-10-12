package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const downloadFolder = "download"

func main() {
	mkdir(downloadFolder)
	for _, link := range os.Args[1:] {
		if numSaved, err := DownloadWebsite(link); err != nil {
			log.Fatalln("error downloading website: ", err.Error())
		} else {
			fmt.Println(numSaved)
		}
	}
}

func DownloadWebsite(link string) (int, error) {
	urlSet := make(map[string]struct{})
	saved := 0

	link = strings.TrimRight(link, "/")
	parsedUrl, err := url.ParseRequestURI(link)
	if err != nil {
		return saved, err
	}

	linkRegexp, err := regexp.Compile("https?://([a-z0-9]+[.])*" + parsedUrl.Host)
	if err != nil {
		return saved, err
	}

	mkdir(downloadFolder + "/" + parsedUrl.Host)

	collector := colly.NewCollector(colly.URLFilters(linkRegexp))

	collector.OnHTML("a[href]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("href"))
		if _, isExist := urlSet[ul]; !isExist {
			urlSet[ul] = struct{}{}
			collector.Visit(ul)
		}
	})

	collector.OnResponse(func(r *colly.Response) {
		reqUrlPath := r.Request.URL.Path
		fullPath := downloadFolder + "/" + parsedUrl.Hostname() + reqUrlPath

		if _, ok := urlSet[fullPath]; ok {
			return
		}

		urlSet[fullPath] = struct{}{}
		if path.Ext(fullPath) == "" {
			mkdir(fullPath)
		} else {
			mkdir(fullPath[:strings.LastIndexByte(fullPath, '/')])
		}

		if path.Ext(reqUrlPath) == "" {
			if fullPath[len(fullPath)-1] != '/' {
				fullPath += "/"
			}
			fullPath += "index.html"
			if _, err := os.Create(fullPath); err != nil {
				fmt.Println("error creating file: ", err.Error())
			}
		}

		if err = r.Save(fullPath); err != nil {
			panic(err)
		}

		fmt.Println("saved: ", parsedUrl.Hostname()+reqUrlPath)
		saved++
	})

	if err = collector.Visit(parsedUrl.String()); err != nil {
		log.Fatalln("error visiting url: ", err.Error())
	}
	collector.Wait()
	return saved, nil
}

func mkdir(folder string) {
	_, err := os.Stat(folder)
	if os.IsNotExist(err) && os.MkdirAll(folder, os.ModePerm) != nil {
		log.Fatalln("error creating directory: ", err.Error())
	}
}
