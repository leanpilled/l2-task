package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getLinks(data []byte, url string) map[string]struct{} {
	links := map[string]struct{}{}
	reader := bytes.NewReader(data)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := tokenizer.Token()

			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if !strings.HasPrefix(link, "https") {
							link = path.Join(url, link)
						}
						links[link] = struct{}{}
					}
				}
			}
		}
	}
}

func writeHTML(filename string, data []byte) error {
	err := os.MkdirAll(filename, os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(filename+"index.html", data, 0644)
}

func wget(dest string, url string, recDepth int) {
	fmt.Println("Downloading ", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writeHTML(path.Join(dest, resp.Request.URL.Path), data)
	if err != nil {
		fmt.Println(err)
		return
	}

	links := getLinks(data, url)

	for link := range links {
		for i := 0; i < recDepth; i++ {
			wget(dest, link, recDepth-i-1)
		}
	}
}

func main() {
	rec := flag.Int("r", 0, "recursionDepth")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		return
	}

	wget(args[0], args[1], *rec)
}
