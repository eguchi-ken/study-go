package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex



// リテラルで定義する場合こんな感じ
// json と少し似ているかも。
var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// ↑と同じ意味。value の型が Vertex とわかっているので一部省略できる。
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

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

	fmt.Println(m2)
}
