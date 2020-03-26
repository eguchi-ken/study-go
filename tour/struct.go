// https://go-tour-jp.appspot.com/moretypes/2

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// Vertex を作る
	v := Vertex{1, 2}

	fmt.Println(v)

	v.X = 4

	fmt.Println(v.X)

	// v のポインタをとる
	p := &v

	// v のポインタから値をとるときはシンプルにドットが使える。
	p.X = 1e9

	fmt.Println((*p).X)
	fmt.Println(p.X)
	fmt.Println(v)

	// いろんな初期化のしかたがある。
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p1  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p1)
}
