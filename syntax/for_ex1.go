package main

import "fmt"

func main() {
	// 0 から 99 まで計100個の数字を順に格納したスライスnums を作成する
	nums := make([]int, 100)
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	fmt.Printf("nums = %v\n", nums)

	// nums を最後の要素から順に出力
	// 出力: 99 98 .... 3 2 1 0
	// for i := len(nums) - 1; i > 0; i-- {
	// 	fmt.Println(nums[i])
	// }

	/*
		// 奇数番目の要素のみを出力
		// 出力: 1 3 5 ... 97 99
	*/

	for i := 0; i < len(nums); i++ {
		if i%3 == 0 {
			fmt.Println(nums[i])
		}
	}
}
