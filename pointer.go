// https://go-tour-jp.appspot.com/moretypes/1

package main

import "fmt"

func main() {
	// 適当な整数を取ってみる。
	i, j := 42, 2701

	// i のアドレスを取り出す(ポインタ)
	p := &i

	// ポインタの指すアドレスに入っている値を表示
	fmt.Println(*p)

	// ポインタの指すアドレスに値をセット
	*p = 21

	// i の値が変化していることを確認する
	fmt.Println(i)

	// j のアドレスを取り出す
	p = &j

	// j の値を 37 で割って、さらに j にセットする
	*p = *p / 37

	// j の最終結果を確認する
	fmt.Println(j)
}
