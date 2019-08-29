// echo コマンドの実装 : http://golang.jp/go_tutorial
package main

import (
  "os"
  "flag"  // コマンドラインオプションのパーサ
)

// -n オプションを用意する。指定した場合は最後に改行を含めないで出力する
var omitNewline = flag.Bool("n", false, "don't print final newline")

// import・var・const あたりはカッコでまとめられる
const (
  Space   = " "
  Newline = "\n"
)
// 以下でも良い
// const Space   = " "
// const Newline = "\n"

func main() {
  // パラメータリストを調べてflag に設定する
  flag.Parse()
  
  // 変数宣言
  var s string = ""  // 型宣言は省略できる
  // s := ""         // こう書いても同じ
  
  // if や for はカッコで囲まなくて良い。while はない
  for i := 0; i < flag.NArg(); i++ {  // ← := 部分が var 変数宣言
    if i > 0 {
      s += Space
    }
    s += flag.Arg(i)
  }
  
  // フラグがなければ改行を付与する
  // アスタリスクはポインタを示す何やらっぽい。不明 https://qiita.com/tmzkysk/items/1b73eaf415fee91aaad3
  if !*omitNewline {
    s += Newline
  }
  
  // 出力する
  os.Stdout.WriteString(s)
}

// エラーを返したい時は os.Exit(1)

// $ go run ./practice-02.go HOGE FUGA
// HOGE FUGA
