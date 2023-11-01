## はてなブログにエディタでMDで書いたブログを飛ばすやつ🚀
kita127さんがはてなブログに公開なさっていた記事を読んで、API経由でブログをポスト出来るという事を知ったので、ブログに書かれていたロジックを参考にGoでざっくり書いてみました。

### 使い方
1.
example.entry.jsonの情報を書き換えて、「entries.json」という名前で保存

2.
```
go run main.go
```
### 参考記事
```
https://kita127.hatenablog.com/entry/2023/05/17/004937
https://developer.hatena.ne.jp/ja/documents/blog/apis/atom/
いずれも閲覧日は11/01/'23
```