package main

import "fmt"
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy, dy)

	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8(i*j)
		}
	}

	return result
}

func buildX(dx int) []uint8 {
	result := make([]uint8, dx, dx)

	for i := 0; i < dx; i++ {
		result[i] = uint8(i)
	}

	return result
}

func main() {
	// 1行作ってみる
	fmt.Println(buildX(10))

	// 二次元配列作ってみる
	fmt.Println(Pic(5, 10))

	// 関数を受け取ってそれを出力するような関数
	pic.Show(Pic)

}

// スライスの長さ、容量、内容を表示する
func printSlice(s []uint8) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
