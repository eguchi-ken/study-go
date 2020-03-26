package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// range 配列(スライス) とすると、インデックスと値が取れる
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// value が要らない場合は省略できる
	for i := range pow {
		fmt.Printf("2**%d = %d\n", i, pow[i])
	}

	// index がいらない場合は _ を使う
	for _,v := range pow {
		fmt.Println(v)
	}


}
