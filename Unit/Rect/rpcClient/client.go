package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 聲明參數結構體，字段首字母大寫
type Params struct {
	// 長和寬
	Width  int
	Height int
}

// 調用服務
func main() {
	// 1.連接遠程RPC服務
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("連接成功...")

	// 2.調用遠程方法
	// 定義接收服務端傳回來的計算結果的變量
	res := 0

	// 求面積
	err = conn.Call("Rect.Area", Params{50, 100}, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("面積: ", res)

	// 求周長
	err = conn.Call("Rect.Perimeter", Params{50, 100}, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("周長: ", res)
}
