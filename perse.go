package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
)

func getEntriesInfo(config *Config) []EntryInfo {
	var titles []string
	var urls []string

	url := fmt.Sprintf(URLTemplate, config.UserID, config.BlogID)
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("HTTPリクエストの送信エラー: %v\n", err)
			return nil
		}
		defer resp.Body.Close()

		//タイトル
		xmlData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("HTTPレスポンスの読み込みエラー: %v\n", err)
			return nil
		}
		titles, urls = parseXML(xmlData, titles, urls)

		//次のページリンク
		nextLink := getNextLink(xmlData)
		url = nextLink
	}

	var info []EntryInfo
	for i := 0; i < len(titles); i++ {
		info = append(info, EntryInfo{Title: titles[i], URL: urls[i]})
	}
	return info
}

func parseXML(xmlData []byte, titles, urls []string) ([]string, []string) {
	doc, err := html.Parse(strings.NewReader(string(xmlData)))
	if err != nil {
		fmt.Printf("HTMLのパースエラー: %v\n", err)
		return titles, urls
	}

	var title, link string
	var inTitle, inLink bool

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			inTitle = true
			return
		} else if n.Type == html.ElementNode && n.Data == "link" {
			for _, attr := range n.Attr {
				if attr.Key == "rel" && attr.Val == "edit" {
					inLink = true
				}
			}
		}

		if inTitle {
			title = n.Data
			inTitle = false
		}
		if inLink {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link = attr.Val
					inLink = false
				}
			}
		}

		if n.FirstChild != nil {
			traverse(n.FirstChild)
		}
		if n.NextSibling != nil {
			traverse(n.NextSibling)
		}
	}

	traverse(doc)

	if title != "" && link != "" {
		titles = append(titles, title)
		urls = append(urls, link)
	}

	return titles, urls
}

func getNextLink(xmlData []byte) string {
	doc, err := html.Parse(strings.NewReader(string(xmlData)))
	if err != nil {
		fmt.Printf("HTMLのパースエラー: %v\n", err)
		return ""
	}

	var nextLink string
	var inLink bool

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, attr := range n.Attr {
				if attr.Key == "rel" && attr.Val == "next" {
					inLink = true
				}
			}
		}

		if inLink {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					nextLink = attr.Val
					inLink = false
				}
			}
		}

		if n.FirstChild != nil {
			traverse(n.FirstChild)
		}
		if n.NextSibling != nil {
			traverse(n.NextSibling)
		}
	}

	traverse(doc)

	return nextLink
}
