package main

import (
	"fmt"
	"io/ioutil"
)

const URLTemplate = "https://blog.hatena.ne.jp/%s/%s/atom/entry"

type Config struct {
	UserID  string
	BlogID  string
	APIKey  string
	Entries []struct {
		Title   string
		SrcPath string
	}
}

type EntryInfo struct {
	Title string
	URL   string
}

func main() {
	config := readJSONFile("entries.json")
	if config == nil {
		fmt.Println("entries.jsonのパースに失敗")
		return
	}

	info := getEntriesInfo(config)

	for _, entry := range config.Entries {
		entryURL := fetchEntryURL(entry.Title, info)
		contents, err := ioutil.ReadFile(entry.SrcPath)
		if err != nil {
			fmt.Println("ファイルの読み込みエラー:", err)
			return
		}
		if entryURL != "" {
			update(entry, config, entryURL, string(contents))
			fmt.Println("記事を更新")
		} else {
			create(entry, config, string(contents))
			fmt.Println("記事を新規作成")
		}
	}
}
