package main

import (
  "./practice-03"  // ココで「ディレクトリ」を読み込む。ディレクトリ配下のファイルの内、大文字で宣言した型などが利用できる
  "fmt"
  "os"
)

func main() {
  // 存在するファイルと存在しないファイル・切り替えて動かしてみる
  // var path = "/PATH/TO/practice-01.go"
  var path = "/does/not/exist"
  
  // ↓以下で指定する 'practiceFile' 部分は、import したディレクトリ配下の .go ファイルで、package 部分で宣言したパッケージ名を指定する
  file, err := practiceFile.Open(path, 0, 0)
  if file == nil {
    fmt.Printf("can't open file; err=%s\n", err)
    os.Exit(1)
  } else {
    // ↑ ココは改行して書くとエラーになる http://sugilog.hatenablog.com/entry/2014/12/29/000855
    fmt.Printf("OK")
  }
}
