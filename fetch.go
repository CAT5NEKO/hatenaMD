package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func fetchEntryURL(title string, info []EntryInfo) string {
	for _, i := range info {
		if i.Title == title {
			return i.URL
		}
	}
	return ""
}

func readJSONFile(filePath string) *Config {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("JSONファイルの読み込みエラー: %v\n", err)
		return nil
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Printf("JSONファイルのパースエラー: %v\n", err)
		return nil
	}

	return &config
}
