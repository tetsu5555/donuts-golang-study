package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// catコマンド
// - コマンドライン引数でファイル名を指定して実行するとそのファイル内容を出力
// - 複数のファイルが与えられた時は、その順番にファイル内容を出力

// - コマンドラインオプション -n が指定された時は、各行の先頭に行数を表示する

func main() {
	file := os.Args[2]
	fmt.Println(file)

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	str := string(raw)

	fmt.Println(str)
}
