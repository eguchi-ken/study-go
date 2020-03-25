package main

import (
	"fmt"
	"math"
)

// 次の平方根の候補を見つける
func find_next(z, x float64) float64 {
	return z - ((z*z - x) / (2 * z))
}

func main() {
	var z float64 = 1
	var base float64 = 2

	// 10回候補を計算して精度をあげる
	for i := 0; i<10; i++ {
		z = find_next(z, base)
		fmt.Println(z)
	}

	// 比較用にライブラリの場合も出す
	fmt.Println(math.Sqrt(2))
}
