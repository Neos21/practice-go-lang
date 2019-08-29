// ファイル読み書き : http://go.shibu.jp/go_tutorial.html#i-o-package

package practiceFile

import (
  // "os"  // os.Error のために使っていたが、使用箇所がなくなったので消す
  "syscall"
)

// 型定義。大文字で始めた型定義や定数などは、パッケージを利用する側から参照できるようになる。エクスポートした状態
type File struct {
  fd   int     // ファイル記述子番号
  name string  // ファイルを開く時の名前
}

// File 型のオブジェクトを作るためのファクトリ関数
func newFile(fd int, name string) *File {
  if fd < 0 {
    return nil
  }
  return &File{fd, name}
}

// 以下のように愚直に作っても別に良い
func newFile2(fd int, name string) *File {
  var myFile = new(File)
  myFile.fd = fd
  myFile.name = name
  return myFile
}

// file と err を返す関数として宣言する
// perm は int じゃダメで「cannot use perm (type int) as type uint32 in argument to syscall.Open」エラーになった。uint32 にした
// err の型は os.Error にしたが、なくなってる様子。「error」にした <http://shinriyo.hateblo.jp/entry/2014/09/23/Web%E3%82%92Golang%E3%81%A7%E3%83%8F%E3%83%9E%E3%81%A3%E3%81%9F>
func Open(name string, mode int, perm uint32) (file *File, err error) {
  // syscall.Open() は複数の値を返すので以下のように受け取る。両方 int 型
  r, e := syscall.Open(name, mode, perm)
  
  // os ライブラリを使って、Unix 系のエラー番号から os.Errno 型のデータに直す
  // err が os.Error でなくなったために、消しておくことにする
  // if e != 0 {
  //   err = os.Errno(e)
  // }
  err = e
  
  return newFile(r, name), err
}
