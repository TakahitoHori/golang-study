package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// 英文章の出現単語数カウント
func WordCount(s string) map[string]int {
	wordCountTable := make(map[string]int) // 単語カウント用テーブルの作成
	// rangeを使用しforeachもどき strings.Fields()で単語ごとに区切る
	for _, v := range strings.Fields(s) {
		wordCountTable[v]++ // 単語をインデックスとし、インクリメント
	}
	return wordCountTable
}

func main() {
	wc.Test(WordCount)
}
