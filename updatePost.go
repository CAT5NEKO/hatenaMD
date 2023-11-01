package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"time"
)

func update(entry struct{ Title, SrcPath string }, config *Config, url, contents string) {
	escaped := html.EscapeString(contents)

	// 記事更新するときにやっていきする更新XMLデータを作成
	xmlData := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	<entry xmlns="http://www.w3.org/2005/Atom">
	  <title>%s</title>
	  <content>%s</content>
	  <updated>%s</updated>
	</entry>`, entry.Title, escaped, time.Now().Format(time.RFC3339)))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(xmlData))
	if err != nil {
		fmt.Printf("HTTPリクエストの作成エラー: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(config.UserID, config.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTPリクエストの送信エラー: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		fmt.Printf("HTTPステータスコード %d\n", resp.StatusCode)
		return
	}
}

func create(entry struct{ Title, SrcPath string }, config *Config, contents string) {
	url := fmt.Sprintf(URLTemplate, config.UserID, config.BlogID)
	escaped := html.EscapeString(contents)

	//POSTリクエストを投げる
	xmlData := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	<entry xmlns="http://www.w3.org/2005/Atom">
	  <title>%s</title>
	  <content>%s</content>
	  <updated>%s</updated>
	</entry>`, entry.Title, escaped, time.Now().Format(time.RFC3339)))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlData))
	if err != nil {
		fmt.Printf("HTTPリクエストの作成エラー: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(config.UserID, config.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTPリクエストの送信エラー: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("HTTPステータスコード %d\n", resp.StatusCode)
		return
	}
}
