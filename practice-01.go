/* 
 * https://news.mynavi.jp/article/gogogo-1/
 */

// パッケージ名 : 付ける
package main

// fmt パッケージを取り込む
import "fmt"

// main 関数を定義する
func main() {
  // 文字列を出力する
  fmt.Printf("hello, world!")
}

// - セミコロンは不要
// - `$ go build ./hoge.go` でコンパイル (2MB 程度になった) → `$ ./hoge` で実行する
// - `$ go run ./hoge.go` でコンパイルせず直接実行できる <http://nununu.hatenablog.jp/entry/2016/09/20/210000>