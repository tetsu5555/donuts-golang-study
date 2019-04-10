package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// JSON Formatter
// - コマンドラインオプションで与えられたファイルの内容をJSONとして解釈,
//   pretty printとして出力
// - コマンドラインオプションでインデントレベル(タブの個数)を指定できるようにする

// 追加仕様
// - コマンドラインオプション -mode によってminify出力できるようにする

type Test struct {
	// 小文字だと外からアクセスできないから大文字にする
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	////////////////////////////////////
	var mode string
	flag.StringVar(&mode, "mode", "", "mode")
	flag.Parse()

	// fmt.Printf("mode = %v\n", mode)
	// fmt.Printf("args = %v\n", flag.Args())
	////////////////////////////////////

	var file string
	if mode == "min" {
		file = os.Args[2]
	} else {
		file = os.Args[1]
	}

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc []Test
	json.Unmarshal(raw, &fc)

	if mode == "min" {
		PpMin(fc)
	} else {
		Pp(fc)
	}
}

func Pp(a []Test) {
	keys := []string{"id", "name", "price"}

	fmt.Println("[")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%10v\n", "{")
		attributes := []interface{}{a[i].Id, a[i].Name, a[i].Price}
		for i, d := range attributes {
			fmt.Printf("%20v : %v", keys[i], d)
			if i < len(attributes)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Printf("%10v", "}")
		if i < len(a)-1 {
			fmt.Println(",")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("]")
}

func PpMin(a []Test) {
	keys := []string{"id", "name", "price"}

	fmt.Printf("%v", "[")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v", "{")
		attributes := []interface{}{a[i].Id, a[i].Name, a[i].Price}
		for j, d := range attributes {
			fmt.Printf("%v : %v", keys[j], d)
			if j < len(attributes)-1 {
				fmt.Printf("%v", ",")
			} else {
				fmt.Printf("%v", "")
			}
		}
		fmt.Printf("%v", "}")
		if i < len(attributes)-2 {
			fmt.Printf("%v", ",")
		} else {
			fmt.Printf("%v", "")
		}
	}
	fmt.Println("]")
}
