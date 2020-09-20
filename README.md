# go_kyopro

Go言語で競技プログラミングを行うときのライブラリ、および提出前に一つのファイルにまとめるコマンドの作成。

## 準備
**macOS以外で動くかは不明**

コマンドでは、`goimports`を使用します。
```
go get golang.org/x/tools/cmd/goimports
```
を実行してください。

## 使い方
`make cmd`で`bin/go-kyopro`が作成されます。main関数を含むファイルを実行すると、`lib/`下にあるファイルなど一つにまとめたgoのソースコードを生成します。
```
# 標準出力する(ファイルに出力したければリダイレクトで)
go-kyopro cmd/sample/main.go

# クリップボードに貼り付ける
go-kyopro -c cmd/sample/main.go
```
