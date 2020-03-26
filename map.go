package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {

	// string -> vertex のマップを作る
	m = make(map[string]Vertex)

	// key: "Bell Labs",
	// value: {40.68433, -74.39967}
	// のキーバリューペアをセットする。
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	// キーから value を取り出す。
	fmt.Println(m["Bell Labs"])
}
