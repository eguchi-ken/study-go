package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 配列リテラル
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)


	// スライス(可変長)
	var slice []int = primes[1:4]
	fmt.Println(slice)

	// 元の配列を変更するとスライスも変更されてる用に見える。
	// なぜならスライスは参照のように振る舞うため。
	primes[1] = 17
	fmt.Println(slice)

	// 逆にスライスの方を変更すると配列の方にも反映される。
	slice[1] = 19
	fmt.Println(primes)

	// スライスリテラル。自動的に長さのない配列が確保されて、
	// そのスライスが取り出されるらしい。
	q := []int{2, 3, 5, 7, 11, 13}
	printSlice(q)

	// スライスの始点を省略できる。その場合は先頭から切り出しになる。
	s0 := primes[:2]
	printSlice(s0)

	// スライスの終点も省略できる。その場合は末尾まで切り出しになる。
	s1 := primes[3:]
	printSlice(s1)

	// スライスの初期値は nil になる。
	var s2 []int
	if s2 == nil {
		printSlice(s2)
	}

	// make を使って、長さと容量を指定したスライスが作れる
	s3 := make([]int, 5, 10)
	printSlice(s3)

	// 容量がたくさんあるので、より大きなスライスを作れる
	s4 := s3[:10]
	s4[5] = 100
	s4[6] = 100
	s4[7] = 100
	printSlice(s4)

	// append は自動的にスライスのサイズを拡張してくれる
	s4 = append(s4,1,2,3,4,5)
	printSlice(s4)
}

// スライスの長さ、容量、内容を表示する
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
