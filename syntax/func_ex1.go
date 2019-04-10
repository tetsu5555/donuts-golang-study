package main

import "fmt"

func main() {
	// 0 から n までの和を計算する関数Sum(n) を作成する
	fmt.Printf("Sum(100) = %d\n", Sum(100))
	// 出力例: "Sum(100) = 5050"
}

func Sum(n int) int {
	sum := 0
	for i := 0; i < n+1; i++ {
		sum += i
	}
	return sum
}
