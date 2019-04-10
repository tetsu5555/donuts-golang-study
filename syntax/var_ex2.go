package main

import "fmt"

func main() {
	v1, v2 := 1, 2
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)

	v1, v2 = v2, v1

	// 変数への代入操作でv1, v2の値を入れ替える
	// ここを修正する
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)
}
